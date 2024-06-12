package api

import (
	"fmt"
	"net/http"
	rediscache "web/redis_cache"
	"web/templates"

	"github.com/google/uuid"
)

func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Api test result")
}

func FetchOrderStatus(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	if id == "" { 
		http.Error(w, "No valid uuid was provided", http.StatusBadRequest)
		return
	}
	orderId, err := uuid.Parse(id);
	if err != nil { 
		http.Error(w, "No valid uuid was provided", http.StatusBadRequest)
		return
	}
	orderInfo, err := rediscache.GetOrderInfo(orderId)
	if err  != nil {
		http.Error(w, "No order with this id was found", http.StatusNotFound)
	} else {
		templates.OrderStatus(orderInfo).Render(req.Context(), w);
	}
}
