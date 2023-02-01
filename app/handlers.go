package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	// to send the response back to client, here we are writing hello world into w
	_, err := fmt.Fprintf(w, "Hello World!\n")
	if err != nil {
		log.Fatal("Error while sending the response err: ", err.Error())
	}
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Akash", City: "Mumbai", Zipcode: "421301"},
		{Name: "Prateek", City: "Mumbai", Zipcode: "421301"},
		{"Suresh", "Banglore", "3878347"},
	}

	// We will be sending response as json or xml based on what client has requested through header
	if r.Header.Get("Content-Type") == "application/xml" {
		// encoding to xml
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Fatal("Errro while sending the response, err: ", err.Error())
		}
	} else {
		// ading content-type
		w.Header().Add("Content-Type", "application/json")
		// encoding our struct to json format and writing it to w which will be sent to client
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Fatal("Error while sending the response err: ", err.Error())
		}
	}

}
