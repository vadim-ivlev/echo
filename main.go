package main

import (
	"echo/api"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/text", api.TextHandler)
	mux.HandleFunc("/api/", api.EchoHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
