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
	mux.HandleFunc("GET /api/orders/{id}", api.FetchOrderStatus)
	mux.HandleFunc("GET /api", api.TestEndpoint)
	mux.HandleFunc("GET /orders/{id}", templates.HandleOrderStatus)
	mux.HandleFunc("POST /submitorder", templates.HandleOrderForm)
	rootHandler := createRootPathHandler(templ.Handler(templates.HomePage()), http.NotFoundHandler())
	mux.HandleFunc("/", rootHandler);
	return mux
}

func createRootPathHandler(rootHandler, notFoundHandler http.Handler) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			rootHandler.ServeHTTP(w, r);
		} else {
			notFoundHandler.ServeHTTP(w, r)
		}
	}
}
