package app

import (
	"fmt"
	"github.com/MohammedKamle/banking/domain"
	"github.com/MohammedKamle/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func StartApplication() {
	// sanity check for env variables
	sanityCheck()
	// creating multiplexer
	router := mux.NewRouter()

	//wiring
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// For Stub
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	ch := CustomerHandler{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
	if err != nil {
		log.Fatal("Error while starting the server")
	}
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "mohammed:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Some or all environment variables are missing....")
	}
}
