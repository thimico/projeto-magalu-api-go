package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service/mocks"
)

func TestCreateContract(t *testing.T) {
	mockService := new(mocks.ContractService)
	ctrl := NewContractController(mockService)

	tenantID := primitive.NewObjectID()
	productID := primitive.NewObjectID()
	now := time.Now().Truncate(time.Second)

	input := &model.ContractIn{
		TenantID:     tenantID,
		ProductID:    productID,
		ContractMode: "monthly",
		StartDate:    now,
		Status:       "active",
	}

	expected := &model.ContractOut{
		ID:           primitive.NewObjectID(),
		TenantID:     tenantID,
		ProductID:    productID,
		ContractMode: "monthly",
		StartDate:    now,
		Status:       "active",
	}

	mockService.
		On("Create", mock.Anything, mock.MatchedBy(func(c *model.ContractIn) bool {
			return c.TenantID == input.TenantID &&
				c.ProductID == input.ProductID &&
				c.ContractMode == input.ContractMode &&
				c.Status == input.Status
		})).
		Return(expected, nil)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/contracts", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	ctrl.CreateContract(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var result model.ContractOut
	json.NewDecoder(rr.Body).Decode(&result)

	assert.Equal(t, expected.ID, result.ID)
	assert.Equal(t, expected.TenantID, result.TenantID)
	assert.Equal(t, expected.ProductID, result.ProductID)
	assert.Equal(t, expected.Status, result.Status)
	assert.Equal(t, expected.ContractMode, result.ContractMode)
	mockService.AssertExpectations(t)
}

func TestGetContract(t *testing.T) {
	mockService := new(mocks.ContractService)
	ctrl := NewContractController(mockService)

	id := primitive.NewObjectID()
	tenantID := primitive.NewObjectID()
	productID := primitive.NewObjectID()
	startDate := time.Now()
	endDate := startDate.Add(24 * time.Hour)

	expected := &model.ContractOut{
		ID:           id,
		TenantID:     tenantID,
		ProductID:    productID,
		ContractMode: "test-mode",
		StartDate:    startDate,
		EndDate:      &endDate,
		Status:       "active",
	}

	mockService.On("GetByID", mock.Anything, id.Hex()).Return(expected, nil)

	req := httptest.NewRequest(http.MethodGet, "/contracts/"+id.Hex(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": id.Hex()})
	rr := httptest.NewRecorder()

	ctrl.GetContract(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result model.ContractOut
	json.NewDecoder(rr.Body).Decode(&result)

	assert.Equal(t, expected.ID, result.ID)
	assert.Equal(t, expected.TenantID, result.TenantID)
	assert.Equal(t, expected.ProductID, result.ProductID)
	assert.Equal(t, expected.ContractMode, result.ContractMode)
	assert.Equal(t, expected.Status, result.Status)
}
