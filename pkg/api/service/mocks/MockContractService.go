package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"projeto-magalu-api-go/pkg/api/model"
)

type ContractService struct {
	mock.Mock
}

func (_m *ContractService) Create(ctx context.Context, contract *model.ContractIn) (*model.ContractOut, error) {
	args := _m.Called(ctx, contract)
	return args.Get(0).(*model.ContractOut), args.Error(1)
}

func (m *ContractService) GetByID(ctx context.Context, id string) (*model.ContractOut, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.ContractOut), args.Error(1)
}
