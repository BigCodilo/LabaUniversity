package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/events", EventsHandler)
	r.HandleFunc("/venues", VenuesHandler)
	r.HandleFunc("/reservant/{Id}", ReservantHandler)
	r.HandleFunc("/customers", CustomersHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	http.Handle("/", r)
	http.ListenAndServe(":1111", nil)
}
