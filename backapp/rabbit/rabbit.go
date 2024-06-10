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
var orderChannel chan OrderUpdateChannelPayload = make(chan OrderUpdateChannelPayload);

const TIMEOUT time.Duration = 5;
const RETRIES int = 5;
const SEND_TIMEOUT time.Duration = 10;

type OrderUpdateChannelPayload struct {
	OrderType	string
	Payload		[]byte 
}

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
			go handleUpdateSend();
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
	// Receive messages from web app
	channel.ExchangeDeclare("translation", "direct", false, false, false, false, nil);
	channel.QueueDeclare("orders", false, false, false, false, nil);
	channel.QueueBind("orders", "translation_orders", "translation", false, nil);
	// Send out messages to topic exchange
	channel.ExchangeDeclare("updates", "topic", false, false, false, false, nil);
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
		go handler.HandleMessage(&msg, &orderChannel)
	}
	<-forever
}

func handleUpdateSend() {
	if orderChannel == nil {
		panic("handleUpdateSend: orderChannel is nil")
	}
	for sendData := range orderChannel {
		ensureChannelHealth();
		routingKey := fmt.Sprintf("order.%s", sendData.OrderType);
		channel.Publish("updates", routingKey, false, false, amqp.Publishing{Body: sendData.Payload});
	}
}
