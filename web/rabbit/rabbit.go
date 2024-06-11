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

const updateQueueName = "updates"

var OrderUpdateChannel chan *messages.OrderStatusInfo = make(chan *messages.OrderStatusInfo);

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
			go handleOrderUpdate();
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
	// Sending new translation info
	channel.ExchangeDeclare("translation", "direct", false, false, false, false, nil);
	channel.QueueDeclare("orders", false, false, false, false, nil);
	channel.QueueBind("orders", "translation_orders", "translation", false, nil);
	// Order updates info
	channel.ExchangeDeclare("updates", "topic", false, false, false, false, nil);
	channel.QueueDeclare(updateQueueName, false, false, true, false, nil);
	channel.QueueBind(updateQueueName, "order.*", "updates", false, nil)
	//Order updates requests
	channel.QueueDeclare("info_requests", false, false, false, false, nil);
	channel.QueueBind("info_requests", "info_requests", "translation", false, nil)
	return nil;
}

func ensureChannelHealth() error {
	if channel == nil || channel.IsClosed() {
		return establishChannel();
	}
	return nil;
}

func handleOrderUpdate() {
	// TODO - actually do the error handling/manual ack logic
	msgs, _ := channel.Consume(
                updateQueueName, // queue
                "webapp",     // consumer
                true,   // auto ack
                true,  // exclusive
                false,  // no local
                false,  // no wait
                nil,    // args
        );
	for msg := range msgs {
		var info messages.OrderStatusInfo;
		err := proto.Unmarshal(msg.Body, &info);
		if err != nil {
			log.Println("ERROR There was an error while serializing proto")
		}
		log.Printf("INFO Received order update: %s\n", info.Id)
		OrderUpdateChannel<-&info
	}
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

func SendOrderUpdateRequest(data *messages.GetOrder) error {
	ensureChannelHealth();
	binaryData, err := proto.Marshal(data);
	if err != nil {
		log.Println("ERROR: rabbit.SendOrderUpdateRequest - error marshalling protobuf")
		return err;
	}
	ctx, cancel := context.WithTimeout(context.Background(), SEND_TIMEOUT)
	defer cancel();
	channel.PublishWithContext(ctx, "translation", "info_requests", false, false, amqp.Publishing{Body: binaryData})
	return nil;
}
