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
	mux.HandleFunc("POST /submitorder", templates.HandleOrderForm)
	mux.Handle("/home", templ.Handler(templates.HomePage()))
	rh := http.RedirectHandler("/home", 307)
	mux.Handle("/", rh);
	return mux
}
