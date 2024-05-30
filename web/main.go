package main

import (
	"fmt"
	"net/http"
	"web/templates"

	"github.com/a-h/templ"
)

func main() {
	println("Starting server");
	hello := templates.HomePage("guy");
	http.Handle("/", templ.Handler(hello));
	fmt.Println("Listening on :2137");
	http.ListenAndServe(":2137", nil);
}
