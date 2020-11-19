package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Person es la estructura de la informacion que va a tener la persona
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//Address es la estructura de la informacion de la direccion
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

//GetPeopleEndpoint devuel al servidor toda la informacion de las personas
func GetPeopleEndpoint(a http.ResponseWriter, request *http.Request) {
	json.NewEncoder(a).Encode(people)
}

//GetPersonEndpoint devuelve la informacion de una persona pedida por el servidor
func GetPersonEndpoint(a http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(a).Encode(item)
			return
		}
	}
	json.NewEncoder(a).Encode(&Person{})

}
