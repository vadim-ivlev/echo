package main

import (
	"echo/api"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.EchoHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
