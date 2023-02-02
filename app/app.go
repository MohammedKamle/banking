package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartApplication() {
	// defining our own multiplexer
	//mux := http.NewServeMux()
	r := mux.NewRouter()

	// define routes, passing reference to handler function
	r.HandleFunc("/greet", greet).Methods(http.MethodGet)
	r.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	// [0-9]+ is a regex which will constraint our api to work only when request path contains integers
	r.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	r.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// starting server
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal("Error while starting the server")
	}
}
