package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service/mocks"
)

func TestCreatePulse(t *testing.T) {
	mockService := new(mocks.PulseService)
	controller := NewPulseController(mockService)

	pulseIn := &model.PulseIn{
		Tenant:     "tenant1",
		ProductSKU: "sku1",
		UsedAmount: 100.0,
		UseUnity:   "GB",
	}

	pulse := pulseIn.ToPulse()

	mockService.On("Create", mock.Anything, pulseIn).Return(pulse, nil)

	body, _ := json.Marshal(pulseIn)
	req, err := http.NewRequest("POST", "/pulse", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreatePulse)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}
