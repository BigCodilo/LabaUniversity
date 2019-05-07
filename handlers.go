package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/MPROJECTS/laba/model"
	"github.com/MPROJECTS/laba/model/datamodel"
	"github.com/gorilla/mux"
)

func ClearIPMap() {
	for {
		//fmt.Println(IPs)
		IPs = make(map[string]int)
		time.Sleep(5 * time.Second)
	}
}

func ConsoleLogger(IPAndPort string, hadnler string, w http.ResponseWriter) bool {
	IP := IPAndPort[:len(IPAndPort)-6]

	valIP, ok := IPs[IP]
	if ok {
		valIP++
		IPs[IP] = valIP
	} else {
		IPs[IP] = 0
	}

	if valIP > 10 {
		w.WriteHeader(404)
		return true
	}

	fmt.Println("Someone connect to ", hadnler, " from ", IP, " with second count ", IPs[IP])
	return false
}

//EventsHandler - 1
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	// if coockie, _ := r.Cookie("Token"); coockie == nil {
	// 	w.WriteHeader(404)
	// 	return
	// }
	if ConsoleLogger(r.RemoteAddr, r.RequestURI, w) {
		return
	}
	tmpl, _ := template.ParseFiles("eventsHtml.html")

	db := &model.DataBase{}
	db.Open()
	defer db.Close()
	events := db.GetEvents()

	type EventVenue struct {
		Event datamodel.Event
		Venue datamodel.Venue
	}

	eventsvenues := []EventVenue{}
	for i := 0; i < len(events); i++ {
		eventvenue := EventVenue{
			events[i],
			db.GetVenueByID(events[i].IDVenue),
		}
		eventsvenues = append(eventsvenues, eventvenue)
	}

	tmpl.ExecuteTemplate(w, "events", eventsvenues)

}

//VenuesHandler - 2
func VenuesHandler(w http.ResponseWriter, r *http.Request) {
	// if coockie, _ := r.Cookie("Token"); coockie == nil {
	// 	w.WriteHeader(404)
	// 	return
	// }
	if ConsoleLogger(r.RemoteAddr, r.RequestURI, w) {
		return
	}
	tmpl, _ := template.ParseFiles("venuesHtml.html")

	db := &model.DataBase{}
	db.Open()
	venues := db.GetVenues()
	defer db.Close()

	tmpl.ExecuteTemplate(w, "venues", venues)
}

//ReservantHandler - 3
func ReservantHandler(w http.ResponseWriter, r *http.Request) {
	// if coockie, _ := r.Cookie("Token"); coockie == nil {
	// 	w.WriteHeader(404)
	// 	return
	// }
	if ConsoleLogger(r.RemoteAddr, r.RequestURI, w) {
		return
	}
	tmpl, _ := template.ParseFiles("reservantHtml.html")
	vars := mux.Vars(r)
	IDCustomer, _ := strconv.Atoi(vars["Id"])

	db := &model.DataBase{}
	db.Open()
	defer db.Close()

	customerM := db.GetCustomerByID(IDCustomer)
	reservantsM := db.GetReservantByIDCustomer(IDCustomer)

	//struct with customer which has bought a ticket, and with event for this ticket

	events := []datamodel.Event{}
	for i := 0; i < len(reservantsM); i++ {
		eventM := db.GetEventByID(reservantsM[i].IDEvent)
		events = append(events, eventM)
	}

	type ReservantCustomerEvent struct {
		Events   []datamodel.Event
		Customer datamodel.Customer
	}

	reservantscustomersevents := ReservantCustomerEvent{
		events,
		customerM,
	}

	tmpl.ExecuteTemplate(w, "reservant", reservantscustomersevents)
}

//CustomersHandler - 4
func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	// if coockie, _ := r.Cookie("Token"); coockie == nil {
	// 	w.WriteHeader(404)
	// 	return
	// }
	if ConsoleLogger(r.RemoteAddr, r.RequestURI, w) {
		return
	}
	tmpl, _ := template.ParseFiles("customersHtml.html")

	db := &model.DataBase{}
	db.Open()
	customers := db.GetCustomers()
	defer db.Close()

	tmpl.ExecuteTemplate(w, "customers", customers)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// if coockie, _ := r.Cookie("Token"); coockie == nil {
	// 	w.WriteHeader(404)
	// 	return
	//	}
	if ConsoleLogger(r.RemoteAddr, r.RequestURI, w) {
		return
	}
	w.Write([]byte("артем и дима пидарасы"))
}

func SetTokenHandler(w http.ResponseWriter, r *http.Request) {
	hour, min, _ := time.Now().Clock()
	token := strconv.Itoa(hour * min)
	cookie := http.Cookie{
		Name:  "Token",
		Value: token,
	}
	http.SetCookie(w, &cookie)
}
