package main

import (
	"fmt"
	"log"
	"net/http"

	gopher "github.com/mateocarranza/Rest_Api/src/Myproyect/pkg"

	"./pkg/server"
)

func main() {
	s := server.New()

	gopher.SetPeople("1", "mateo", "carranza", "cordoba", "cordoba")

	fmt.Println("Server abierto en http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}
