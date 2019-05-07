package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var IPs map[string]int

func main() {

	go ClearIPMap()

	r := mux.NewRouter()

	r.HandleFunc("/events", EventsHandler)
	r.HandleFunc("/venues", VenuesHandler)
	r.HandleFunc("/reservant/{Id}", ReservantHandler)
	r.HandleFunc("/customers", CustomersHandler)
	r.HandleFunc("/setToken", SetTokenHandler)
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("", IndexHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	http.Handle("/", r)
	http.ListenAndServe(":1111", nil)
}
