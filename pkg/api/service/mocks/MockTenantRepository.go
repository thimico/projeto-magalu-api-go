package mocks

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-magalu-api-go/pkg/api/model"
)

// MockTenantRepository is a mock implementation of the TenantRepository interface
type MockTenantRepository struct {
	tenants map[string]*model.Tenant
}

func NewMockTenantRepository() *MockTenantRepository {
	return &MockTenantRepository{
		tenants: make(map[string]*model.Tenant),
	}
}

func (m *MockTenantRepository) Create(ctx context.Context, tenant *model.Tenant) (*model.Tenant, error) {
	tenant.ID = primitive.NewObjectID()
	m.tenants[tenant.ID.Hex()] = tenant
	return tenant, nil
}

func (m *MockTenantRepository) GetByID(ctx context.Context, id string) (*model.Tenant, error) {
	tenant, exists := m.tenants[id]
	if !exists {
		return nil, errors.New("tenant not found")
	}
	return tenant, nil
}
