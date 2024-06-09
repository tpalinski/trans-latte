package main

import (
	"backapp/db"
	"backapp/rabbit"
	"log"
)

func main() {
	log.Println("Starting backend app");
	log.Println("Initializing postgres connection")
	go db.InitPostgresConnection();
	log.Println("Initializing rabbitmq connection")
	rabbit.IntializeRMQClient(&rabbit.DefaultOrderMessageHandler{});
}
