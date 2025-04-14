package mocks

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-magalu-api-go/pkg/api/model"
)

type MockTenantService struct{}

func (m *MockTenantService) Create(ctx context.Context, in *model.TenantIn) (*model.TenantOut, error) {
	return &model.TenantOut{
		ID:             primitive.NewObjectID(),
		DocumentNumber: in.DocumentNumber,
	}, nil
}

func (m *MockTenantService) GetByID(ctx context.Context, id string) (*model.TenantOut, error) {
	return &model.TenantOut{
		ID:             primitive.NewObjectID(),
		DocumentNumber: "12345678900",
	}, nil
}
