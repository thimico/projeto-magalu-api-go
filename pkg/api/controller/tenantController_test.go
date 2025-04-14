package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projeto-magalu-api-go/pkg/api/controller/mocks"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"projeto-magalu-api-go/pkg/api/model"
)

func TestCreateTenant(t *testing.T) {
	mockService := &mocks.MockTenantService{}
	controller := NewTenantController(mockService)

	tenantIn := &model.TenantIn{
		Name:           "tenantOne",
		DocumentNumber: "12345678900",
	}
	body, _ := json.Marshal(tenantIn)
	req, err := http.NewRequest("POST", "/tenants", bytes.NewBuffer(body))
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.CreateTenant)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var tenantOut model.TenantOut
	err = json.NewDecoder(rr.Body).Decode(&tenantOut)
	assert.NoError(t, err)
	assert.Equal(t, tenantIn.DocumentNumber, tenantOut.DocumentNumber)
}

func TestGetTenant(t *testing.T) {
	mockService := &mocks.MockTenantService{}
	controller := NewTenantController(mockService)

	req, err := http.NewRequest("GET", "/tenants/{id}", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tenants/{id}", controller.GetTenant)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var tenantOut model.TenantOut
	err = json.NewDecoder(rr.Body).Decode(&tenantOut)
	assert.NoError(t, err)
	assert.Equal(t, "12345678900", tenantOut.DocumentNumber)
}
