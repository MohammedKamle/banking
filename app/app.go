package app

import (
	"fmt"
	"github.com/MohammedKamle/banking/domain"
	"github.com/MohammedKamle/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func StartApplication() {
	// sanity check for env variables
	sanityCheck()
	// creating multiplexer
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
	if err != nil {
		log.Fatal("Error while starting the server")
	}
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Some or all environment variables are missing....")
	}
}
