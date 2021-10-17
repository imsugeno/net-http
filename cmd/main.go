package main

import (
	"net/http"

	"github.com/imsugeno/net-http/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/public/hello", handler.Hello)
	http.ListenAndServe(":3000", mux)
}
