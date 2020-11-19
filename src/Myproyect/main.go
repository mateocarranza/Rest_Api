package main

import (
	"fmt"
	"log"
	"net/http"

	"./pkg/server"
)

func main() {
	s := server.New()

	fmt.Println("Server abierto en http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}
