package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
)

type Contract interface {
	Create(ctx context.Context, contract *model.ContractIn) (*model.ContractOut, error)
	GetByID(ctx context.Context, id string) (*model.ContractOut, error)
}

type ContractRepository interface {
	Create(ctx context.Context, contract *model.Contract) (*model.Contract, error)
	GetByID(ctx context.Context, id string) (*model.Contract, error)
}

type contract struct {
	repository ContractRepository
}

func NewContract(repository ContractRepository) Contract {
	return &contract{repository: repository}
}

func (s *contract) Create(ctx context.Context, in *model.ContractIn) (*model.ContractOut, error) {
	contractModel := in.ToContract()
	contract, err := s.repository.Create(ctx, contractModel)
	if err != nil {
		return nil, err
	}
	return contract.ToContractOut(), nil
}

func (s *contract) GetByID(ctx context.Context, id string) (*model.ContractOut, error) {
	contract, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return contract.ToContractOut(), nil
}
