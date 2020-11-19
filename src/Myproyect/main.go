package main

import (
	"log"
	"net/http"

	"./pkg/server"
)

func main() {
	s := server.New()

	people = append(people, Person{ID: "1", FirstName: "Mateo", LastName: "Carranza", Address: &Address{City: "Cordoba", State: "Cordoba"}})
	people = append(people, Person{ID: "2", FirstName: "Carlos", LastName: "Carranza", Address: &Address{City: "Cordoba", State: "Cordoba"}})

	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}
