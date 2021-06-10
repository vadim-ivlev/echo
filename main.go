package main

import (
	"fmt"
	"github.com/vadim-ivlev/echo/api"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Do requests to http://localhost:3000/api/")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", api.EchoHttp)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
