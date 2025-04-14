package controller

import (
	"bytes"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"projeto-magalu-api-go/pkg/api/controller/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"projeto-magalu-api-go/pkg/api/model"
)

func TestCreateTransaction(t *testing.T) {
	mockService := &mocks.MockTransactionService{}
	controller := NewTransactionController(mockService)

	transactionIn := &model.TransactionIn{
		TenantID:        primitive.NewObjectID(),
		OperationTypeID: 4,
		Amount:          100,
	}
	body, _ := json.Marshal(transactionIn)
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.CreateTransaction)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var transactionOut model.TransactionOut
	err = json.NewDecoder(rr.Body).Decode(&transactionOut)
	assert.NoError(t, err)
	assert.Equal(t, transactionIn.Amount, transactionOut.Amount)
}
