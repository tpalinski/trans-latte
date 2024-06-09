package main

import (
	"log"
	"net/http"
	"web/rabbit"
	rediscache "web/redis_cache"
	"web/storage"
)

func main() {
	log.Println("Starting server");
	mux := InitializeRoutes();
	log.Println("Connecting to minio instance")
	go storage.InitializeMinioClient();
	log.Println("Connecting to rabbitmq instance")
	go rabbit.IntializeRMQClient();
	log.Println("Connecting to redis server")
	go rediscache.InitializeRedisConnection();
	log.Println("Listening on :2137");
	http.ListenAndServe(":2137", mux);
}
