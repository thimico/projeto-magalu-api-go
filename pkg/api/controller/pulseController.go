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

type PulseCtrl interface {
	CreatePulse(w http.ResponseWriter, r *http.Request)
	GetCurrentMonthConsumption(w http.ResponseWriter, r *http.Request)
	GetAllResourcesConsumption(w http.ResponseWriter, r *http.Request)
}

type PulseCtrlImpl struct {
	service service.Pulse
}

// swagger:route GET /tenants/{tenant}/sku/{sku}/consumption pulses getCurrentMonthConsumption
// Get current month consumption for a specific SKU
// responses:
//
//	200: description: Successful retrieval of consumption data
//	404: ErrorResponse
func (p *PulseCtrlImpl) GetCurrentMonthConsumption(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	tenant := vars["tenant"]
	sku := vars["sku"]

	consumption, err := p.service.GetCurrentMonthConsumption(context.Background(), tenant, sku)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(consumption)
}

// swagger:route GET /tenants/{tenant}/consumption pulses getAllResourcesConsumption
// Get all resources consumption for a tenant
// responses:
//
//	200: description: Successful retrieval of all resources consumption data
//	404: ErrorResponse
func (p *PulseCtrlImpl) GetAllResourcesConsumption(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	tenant := vars["tenant"]

	consumption, err := p.service.GetAllResourcesConsumption(context.Background(), tenant)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(consumption)
}

func NewPulseController(service service.Pulse) *PulseCtrlImpl {
	return &PulseCtrlImpl{service: service}
}

// swagger:route POST /pulses pulses createPulse
// Create a new pulse
// responses:
//
//	201: Pulse
//	400: ErrorResponse
func (p *PulseCtrlImpl) CreatePulse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var in model.PulseIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	error := ValidatePulse(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	pulse, err := p.service.Create(context.Background(), &in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	encoder.Encode(pulse)
}

func ValidatePulse(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
