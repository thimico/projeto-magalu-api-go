package controller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service"
)

type ContractCtrl interface {
	CreateContract(w http.ResponseWriter, r *http.Request)
	GetContract(w http.ResponseWriter, r *http.Request)
}

type ContractCtrlImpl struct {
	service service.Contract
}

func NewContractController(service service.Contract) *ContractCtrlImpl {
	return &ContractCtrlImpl{service: service}
}

// CreateContract handles the creation of a new contract
// swagger:route POST /contracts createContract
// Create a new contract with the input payload
// responses:
//
//	201: ContractOut
//	400: ErrorResponse
//	422: ErrorResponse
//
// responses:
//
//	201: ContractResponse
func (p *ContractCtrlImpl) CreateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var in model.ContractIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	error := ValidateContract(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	contract, err := p.service.Create(context.Background(), &in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	encoder.Encode(contract)
}

// swagger:route GET /contracts/{contractId} contracts getContract
// Get contract details by contract ID
// responses:
//
//	200: ContractOut
//	404: ErrorResponse
func (p *ContractCtrlImpl) GetContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	contract, err := p.service.GetByID(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contract)
}

func ValidateContract(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
