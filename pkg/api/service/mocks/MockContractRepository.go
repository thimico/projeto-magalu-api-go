package mocks

import (
	"context"
	"errors"
	"projeto-magalu-api-go/pkg/api/model"
)

type MockContractRepository struct {
	contracts map[string]*model.Contract
}

func NewMockContractRepository() *MockContractRepository {
	return &MockContractRepository{
		contracts: make(map[string]*model.Contract),
	}
}

func (m *MockContractRepository) Create(ctx context.Context, contract *model.Contract) (*model.Contract, error) {
	m.contracts[contract.ID.String()] = contract
	return contract, nil
}

func (m *MockContractRepository) GetByID(ctx context.Context, id string) (*model.Contract, error) {
	contract, exists := m.contracts[id]
	if !exists {
		return nil, errors.New("contract not found")
	}
	return contract, nil
}
