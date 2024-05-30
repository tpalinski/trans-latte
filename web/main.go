package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("Starting server");
	mux := InitializeRoutes();
	fmt.Println("Listening on :2137");
	http.ListenAndServe(":2137", mux);
}
