package main

import (
	"backapp/rabbit"
	"log"
)

func main() {
	log.Println("Starting backend app");
	log.Println("Initializing rabbitmq connection")
	rabbit.IntializeRMQClient(&rabbit.DefaultOrderMessageHandler{});
}
