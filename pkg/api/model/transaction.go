package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Transaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"transaction_id"`
	TenantID        primitive.ObjectID `bson:"tenant_id" json:"tenant_id"`
	OperationTypeID int                `bson:"operation_type_id" json:"operation_type_id"`
	Amount          float64            `bson:"amount" json:"amount"`
	EventDate       time.Time          `bson:"event_date" json:"event_date"`
	Balance         float64            `bson:"balance" json:"balance"`
}

type TransactionIn struct {
	TenantID        primitive.ObjectID `json:"tenant_id" validate:"required"`
	OperationTypeID int                `json:"operation_type_id" validate:"required"`
	Amount          float64            `json:"amount" validate:"required"`
}

type TransactionOut struct {
	ID              primitive.ObjectID `json:"transaction_id"`
	TenantID        primitive.ObjectID `json:"tenant_id"`
	OperationTypeID int                `json:"operation_type_id"`
	Amount          float64            `json:"amount"`
	EventDate       time.Time          `json:"event_date"`
	Balance         float64            `json:"balance"`
}

func (t *TransactionIn) ToTransaction() *Transaction {
	return &Transaction{
		TenantID:        t.TenantID,
		OperationTypeID: t.OperationTypeID,
		Amount:          t.Amount,
		EventDate:       time.Now(),
	}
}

func (t *Transaction) ToTransactionOut() *TransactionOut {
	return &TransactionOut{
		ID:              t.ID,
		TenantID:        t.TenantID,
		OperationTypeID: t.OperationTypeID,
		Amount:          t.Amount,
		EventDate:       t.EventDate,
		Balance:         t.Balance,
	}
}
