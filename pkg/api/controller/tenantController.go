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

type TenantCtrl interface {
	CreateTenant(w http.ResponseWriter, r *http.Request)
	GetTenant(w http.ResponseWriter, r *http.Request)
}

type TenantCtrlImpl struct {
	service service.Tenant
}

func NewTenantController(service service.Tenant) *TenantCtrlImpl {
	return &TenantCtrlImpl{service: service}
}

// CreateTenant handles the creation of a new tenant
// swagger:route POST /tenants createTenant
// Create a new tenant with the input payload
// responses:
//   201: TenantOut
//   400: ErrorResponse
//   422: ErrorResponse
// responses:
//
//	201: TenantResponse
func (p *TenantCtrlImpl) CreateTenant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var in model.TenantIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	error := ValidateTenant(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	tenant, err := p.service.Create(context.Background(), &in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	encoder.Encode(tenant)
}

// GetTenant handles fetching the details of a tenant by ID
// swagger:route GET /tenants/{tenantId} tenants getTenant
// Get tenant details by tenant ID
// responses:
//   200: TenantOut
//   404: ErrorResponse
//   400: ErrorResponse
// responses:
//
//	200: TenantResponse
//	404: ErrorResponse
func (p *TenantCtrlImpl) GetTenant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	tenant, err := p.service.GetByID(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenant)
}

func ValidateTenant(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
