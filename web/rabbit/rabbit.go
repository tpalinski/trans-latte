package rabbit

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"web/messages"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

var connection *amqp.Connection;
var channel *amqp.Channel;

const TIMEOUT time.Duration = 5;
const RETRIES int = 5;
const SEND_TIMEOUT time.Duration = 10;

func IntializeRMQClient() {
	rabbitAddress, ok := os.LookupEnv("RABBITMQ_ADDRESS");
	if !ok {
		rabbitAddress = "localhost"
	}
	connectionString := fmt.Sprintf("amqp://guest:guest@%s:5672/", rabbitAddress)
	for range RETRIES {
		conn, err := amqp.Dial(connectionString)
		if err != nil {
			log.Println("Could not connect to rmq, retrying")
			time.Sleep(TIMEOUT * time.Second)
		} else { // Initialize client
			connection = conn;
			defer connection.Close();
			var forever chan struct {};
			err := establishChannel();
			if err != nil {
				continue;
			}
			log.Println("Connected to rmq instance")
			<-forever;
		}
	}
}

func establishChannel() error {
	c, err := connection.Channel();
	if err != nil {
		log.Println("ERROR: Could not create channel")
		return amqp.Error{};
	}
	channel = c;
	channel.ExchangeDeclare("translation", "direct", false, false, false, false, nil);
	channel.QueueDeclare("orders", false, false, false, false, nil);
	channel.QueueBind("orders", "translation_orders", "translation", false, nil);
	return nil;
}

func ensureChannelHealth() error {
	if channel == nil || channel.IsClosed() {
		return establishChannel();
	}
	return nil;
}

func SendOrderInfo(data *messages.NewOrder, outChannel chan error) {
	err := ensureChannelHealth();
	binaryData, err := proto.Marshal(data);
	if err != nil {
		log.Println("ERROR: rabbit.SendOrderInfo - error marshalling protobuf")
		outChannel<-amqp.Error{} ;
		return;
	}
	ctx, cancel := context.WithTimeout(context.Background(), SEND_TIMEOUT)
	defer cancel();
	channel.PublishWithContext(ctx, "translation", "translation_orders", false, false, amqp.Publishing{Body: binaryData, ContentType: "data/binary"})
	outChannel<-nil;
}
