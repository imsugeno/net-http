package main

import (
	"net/http"

	"github.com/imsugeno/net-http/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/public/hello", handler.Hello)
}
