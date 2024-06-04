package main

import (
	"log"
	"net/http"
	"web/rabbit"
	"web/storage"
)

func main() {
	log.Println("Starting server");
	mux := InitializeRoutes();
	log.Println("Connecting to minio instance")
	storage.InitializeMinioClient();
	log.Println("Done!");
	log.Println("Connecting to rabbitmq instance")
	go rabbit.IntializeRMQClient();
	log.Println("Listening on :2137");
	http.ListenAndServe(":2137", mux);
}
