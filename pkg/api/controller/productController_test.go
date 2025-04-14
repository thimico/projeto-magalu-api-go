package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service/mocks"
)

func TestCreateProduct(t *testing.T) {
	mockService := new(mocks.ProductService)
	controller := NewProductController(mockService)

	productIn := &model.ProductIn{
		Name:       "Product1",
		Price:      10.0,
		MetricUnit: "GB",
	}

	id := primitive.NewObjectID()
	productOut := &model.ProductOut{
		ID:         id,
		Name:       "Product1",
		Price:      10.0,
		MetricUnit: "GB",
	}

	mockService.On("Create", mock.Anything, productIn).Return(productOut, nil)

	body, _ := json.Marshal(productIn)
	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateProduct)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response model.ProductOut
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, productOut.ID, response.ID)
	assert.Equal(t, productOut.Name, response.Name)
	assert.Equal(t, productOut.Price, response.Price)
	assert.Equal(t, productOut.MetricUnit, response.MetricUnit)

	mockService.AssertExpectations(t)
}

func TestGetProduct(t *testing.T) {
	mockService := new(mocks.ProductService)
	controller := NewProductController(mockService)

	product := &model.ProductOut{}
	mockService.On("GetByID", mock.Anything, "1").Return(product, nil)

	req, err := http.NewRequest("GET", "/product/1", nil)
	assert.NoError(t, err)

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetProduct)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}
