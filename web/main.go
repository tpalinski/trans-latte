package main

import (
	"log"
	"net/http"
	"web/storage"
)

func main() {
	log.Println("Starting server");
	mux := InitializeRoutes();
	log.Println("Connecting to minio instance")
	storage.InitializeMinioClient();
	log.Println("Listening on :2137");
	http.ListenAndServe(":2137", mux);
}
