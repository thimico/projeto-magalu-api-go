package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service/mocks"
	"testing"
)

func TestContract_Create(t *testing.T) {
	repo := mocks.NewMockContractRepository()
	service := NewContract(repo)

	contractIn := &model.ContractIn{}
	_, err := service.Create(context.Background(), contractIn)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestContract_GetByID(t *testing.T) {
	repo := mocks.NewMockContractRepository()
	service := NewContract(repo)

	_, err := service.GetByID(context.Background(), "contract1")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
