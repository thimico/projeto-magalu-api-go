package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/service/mocks"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-magalu-api-go/pkg/api/model"
)

func TestTenant_Create(t *testing.T) {
	mockRepo := mocks.NewMockTenantRepository()
	tenantService := NewTenant(mockRepo)

	ctx := context.Background()
	tenantIn := &model.TenantIn{
		DocumentNumber: "123456789",
	}

	tenantOut, err := tenantService.Create(ctx, tenantIn)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tenantOut.DocumentNumber != tenantIn.DocumentNumber {
		t.Errorf("expected document number %v, got %v", tenantIn.DocumentNumber, tenantOut.DocumentNumber)
	}

	if tenantOut.ID.IsZero() {
		t.Errorf("expected non-zero tenant ID")
	}
}

func TestTenant_GetByID(t *testing.T) {
	mockRepo := mocks.NewMockTenantRepository()
	tenantService := NewTenant(mockRepo)

	ctx := context.Background()
	tenantIn := &model.TenantIn{
		DocumentNumber: "123456789",
	}

	createdTenant, err := tenantService.Create(ctx, tenantIn)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	retrievedTenant, err := tenantService.GetByID(ctx, createdTenant.ID.Hex())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(createdTenant, retrievedTenant) {
		t.Errorf("expected tenant %v, got %v", createdTenant, retrievedTenant)
	}
}

func TestTenant_GetByID_NotFound(t *testing.T) {
	mockRepo := mocks.NewMockTenantRepository()
	tenantService := NewTenant(mockRepo)

	ctx := context.Background()
	nonExistentID := primitive.NewObjectID().Hex()

	_, err := tenantService.GetByID(ctx, nonExistentID)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := "tenant not found"
	if err.Error() != expectedError {
		t.Errorf("expected error %v, got %v", expectedError, err.Error())
	}
}
