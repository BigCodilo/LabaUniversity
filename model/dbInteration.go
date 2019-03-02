package model

import (
	"database/sql"

	"github.com/MPROJECTS/laba/model/datamodel"
	_ "github.com/go-sql-driver/mysql" //justifying
)

//DataBase - here the connection
type DataBase struct {
	Connection *sql.DB
}

//Open - open connection with database
func (db *DataBase) Open() {
	db.Connection, _ = sql.Open("mysql", "root:root@/laba")
}

//GetCustomers - get all customers from database
func (db *DataBase) GetCustomers() []datamodel.Customer {
	rows, _ := db.Connection.Query("SELECT * FROM laba.customers")
	defer rows.Close()
	customers := []datamodel.Customer{}
	for rows.Next() {
		customer := datamodel.Customer{}
		rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.SecondName,
			&customer.Age,
			&customer.Passport,
			&customer.Addres,
		)
		customers = append(customers, customer)
	}
	return customers
}

//GetReservantByIDCustomer - get a reservant from database when IDCustomer = id
func (db *DataBase) GetReservantByIDCustomer(id int) []datamodel.Reservant {
	row, _ := db.Connection.Query("SELECT * FROM laba.reservants WHERE idCustomer = ?", id)
	defer row.Close()
	reservants := []datamodel.Reservant{}
	for row.Next() {
		reservant := datamodel.Reservant{}
		row.Scan(
			&reservant.ID,
			&reservant.IDEvent,
			&reservant.IDCustomer,
		)
		reservants = append(reservants, reservant)
	}
	return reservants
}

//GetEvents - get all events from database
func (db *DataBase) GetEvents() []datamodel.Event {
	rows, _ := db.Connection.Query("SELECT * FROM laba.events")
	defer rows.Close()
	events := []datamodel.Event{}
	for rows.Next() {
		event := datamodel.Event{}
		rows.Scan(
			&event.ID,
			&event.IDVenue,
			&event.Name,
			&event.Time,
		)
		events = append(events, event)
	}
	return events
}

//GetVenues - get all venues from database
func (db *DataBase) GetVenues() []datamodel.Venue {
	rows, _ := db.Connection.Query("SELECT * FROM laba.venues")
	defer rows.Close()
	venues := []datamodel.Venue{}
	for rows.Next() {
		venue := datamodel.Venue{}
		rows.Scan(
			&venue.ID,
			&venue.Name,
			&venue.Type,
			&venue.Addres,
			&venue.ImageURL,
		)
		venues = append(venues, venue)
	}
	return venues
}

//GetCustomerByID - get customer by him id from database
func (db *DataBase) GetCustomerByID(id int) datamodel.Customer {
	row, _ := db.Connection.Query("SELECT * FROM laba.customers WHERE id = ?", id)
	defer row.Close()
	row.Next()
	customer := datamodel.Customer{}
	row.Scan(
		&customer.ID,
		&customer.Name,
		&customer.SecondName,
		&customer.Age,
		&customer.Passport,
		&customer.Addres,
	)
	return customer
}

//GetEventByID - get event by id from database
func (db *DataBase) GetEventByID(id int) datamodel.Event {
	row, _ := db.Connection.Query("SELECT * FROM laba.events WHERE id = ?", id)
	defer row.Close()
	row.Next()
	event := datamodel.Event{}
	row.Scan(
		&event.ID,
		&event.IDVenue,
		&event.Name,
		&event.Time,
	)
	return event
}

//GetVenueByID - get venue by id from database
func (db *DataBase) GetVenueByID(id int) datamodel.Venue {
	row, _ := db.Connection.Query("SELECT * FROM laba.venues WHERE id = ?", id)
	defer row.Close()
	row.Next()
	venue := datamodel.Venue{}
	row.Scan(
		&venue.ID,
		&venue.Name,
		&venue.Type,
		&venue.Addres,
		&venue.ImageURL,
	)
	return venue
}

//Close - close the connection with database
func (db *DataBase) Close() {
	db.Connection.Close()
}
