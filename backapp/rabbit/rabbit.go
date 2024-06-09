package rabbit

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var connection *amqp.Connection;
var channel *amqp.Channel;

const TIMEOUT time.Duration = 5;
const RETRIES int = 5;
const SEND_TIMEOUT time.Duration = 10;

func IntializeRMQClient(orderHandler OrderMessageHandler) {
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
			go handleIncomingOrders(orderHandler)
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

func handleIncomingOrders(handler OrderMessageHandler) {
	orderChan, err := connection.Channel();
	if err != nil {
		log.Printf("ERROR: Could not open order handling channel")
	}
	defer orderChan.Close();
	ctx := context.Background();
	var forever chan struct{}
	msgs, _ := orderChan.ConsumeWithContext(ctx, "orders", "backapp", true, false, false, false, nil);
	for msg := range msgs {
		go handler.HandleMessage(&msg)
	}
	<-forever
}
