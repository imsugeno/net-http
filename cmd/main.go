package main

import (
	"net/http"

	"net-http/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/public/hello", handler.Hello)
}
