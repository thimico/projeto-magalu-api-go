package mocks

import (
	"context"
	"errors"
	"projeto-magalu-api-go/pkg/api/model"
)

type MockProductRepository struct {
	products map[string]*model.Product
}

func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products: make(map[string]*model.Product),
	}
}

func (m *MockProductRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	m.products[product.ID.String()] = product
	return product, nil
}

func (m *MockProductRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	product, exists := m.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}
