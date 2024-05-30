package api

import (
	"fmt"
	"net/http"
)

func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Api test result")
}
