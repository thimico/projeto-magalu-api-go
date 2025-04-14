package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service/mocks"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	repo := mocks.NewMockProductRepository()
	service := NewProduct(repo)

	productIn := &model.ProductIn{}
	_, err := service.Create(context.Background(), productIn)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestProduct_GetByID(t *testing.T) {
	repo := mocks.NewMockProductRepository()
	service := NewProduct(repo)

	_, err := service.GetByID(context.Background(), "product1")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
