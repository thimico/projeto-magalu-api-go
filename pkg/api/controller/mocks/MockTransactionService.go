package mocks

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-magalu-api-go/pkg/api/model"
)

type MockTransactionService struct{}

func (m *MockTransactionService) Create(ctx context.Context, in *model.TransactionIn) (*model.TransactionOut, error) {
	return &model.TransactionOut{
		ID:       primitive.NewObjectID(),
		TenantID: in.TenantID,
		Amount:   in.Amount,
	}, nil
}

func (m *MockTransactionService) BalancePayment(ctx context.Context, in *model.TransactionIn) (*model.TransactionOut, error) {
	return &model.TransactionOut{
		ID:       primitive.NewObjectID(),
		TenantID: in.TenantID,
		Amount:   in.Amount,
	}, nil
}
