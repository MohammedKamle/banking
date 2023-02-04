package app

import (
	"github.com/MohammedKamle/banking/domain"
	"github.com/MohammedKamle/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartApplication() {
	// creating multiplexer
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("Error while starting the server")
	}
}
