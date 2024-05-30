package main

import (
	"net/http"
	"web/api"
	"web/templates"

	"github.com/a-h/templ"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("GET /api", api.TestEndpoint)
	mux.Handle("/", templ.Handler(templates.HomePage()))
	return mux
}
