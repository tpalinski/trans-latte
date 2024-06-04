package rabbit

import (
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var connection *amqp.Connection;

const TIMEOUT time.Duration = 5;
const RETRIES int = 5;

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
			c, err := connection.Channel();
			if err != nil {
				log.Println("ERROR: Could not create channel")
			}
			c.ExchangeDeclare("translation", "direct", false, false, false, false, nil);
			c.QueueDeclare("orders", false, false, false, false, nil);
			c.QueueBind("orders", "translation_orders", "translation", false, nil);
			log.Println("Connected to rmq instance")
			<-forever;
		}
	}
}
