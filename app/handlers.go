package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/MohammedKamle/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//
	customers, _ := ch.service.GetAllCustomers()
	// We will be sending response as json or xml based on what client has requested through header
	if r.Header.Get("Content-Type") == "application/xml" {
		// encoding to xml
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		// ading content-type
		w.Header().Add("Content-Type", "application/json")
		// encoding our struct to json format and writing it to w which will be sent to client
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
