package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

//Server es la interface nuestro server
type Server interface {
	Router() http.Handler
}

//New Server cre un servidor nuevo y lo retorna
func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/people", GetPeopleEndpoint).Methods("Get")
	r.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("Get")

	a.router = r

	return a
}

func (a *api) Router() http.Handler {

	return a.router
}
