package app

import (
	"encoding/json"
	"github.com/MohammedKamle/banking/dto"
	"github.com/MohammedKamle/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	// We need to fetch the customer_id from request URL and set it NewAccountRequest DTO
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appErr := h.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
