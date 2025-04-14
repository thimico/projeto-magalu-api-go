package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
)

type Transaction interface {
	Create(ctx context.Context, transaction *model.TransactionIn) (*model.TransactionOut, error)
	BalancePayment(ctx context.Context, transaction *model.TransactionIn) (*model.TransactionOut, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
	BalancePayment(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

type transaction struct {
	repository TransactionRepository
}

func NewTransaction(repository TransactionRepository) Transaction {
	return &transaction{repository: repository}
}

func (s *transaction) Create(ctx context.Context, in *model.TransactionIn) (*model.TransactionOut, error) {
	transactionModel := in.ToTransaction()
	transactionModel.Balance = transactionModel.Amount
	transaction, err := s.repository.Create(ctx, transactionModel)
	if err != nil {
		return nil, err
	}
	return transaction.ToTransactionOut(), nil
}

func (s *transaction) BalancePayment(ctx context.Context, in *model.TransactionIn) (*model.TransactionOut, error) {
	transactionModel := in.ToTransaction()
	transaction, err := s.repository.BalancePayment(ctx, transactionModel)
	if err != nil {
		return nil, err
	}
	return transaction.ToTransactionOut(), nil
}
