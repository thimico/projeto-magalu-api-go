package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Contract struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"contract_id"`
	TenantID     primitive.ObjectID `bson:"tenant_id" json:"tenant_id"`
	ProductID    primitive.ObjectID `bson:"product_id" json:"product_id"`
	ContractMode string             `bson:"contract_mode" json:"contract_mode"`
	StartDate    time.Time          `bson:"start_date" json:"start_date"`
	EndDate      *time.Time         `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Status       string             `bson:"status" json:"status"`
}

type ContractIn struct {
	TenantID     primitive.ObjectID `json:"tenant_id" validate:"required"`
	ProductID    primitive.ObjectID `json:"product_id" validate:"required"`
	ContractMode string             `json:"contract_mode" validate:"required,oneof=monthly yearly on_demand"`
	StartDate    time.Time          `json:"start_date" validate:"required"`
	EndDate      *time.Time         `json:"end_date,omitempty"`
	Status       string             `json:"status" validate:"required"`
}

type ContractOut struct {
	ID           primitive.ObjectID `json:"contract_id"`
	TenantID     primitive.ObjectID `json:"tenant_id"`
	ProductID    primitive.ObjectID `json:"product_id"`
	ContractMode string             `json:"contract_mode"`
	StartDate    time.Time          `json:"start_date"`
	EndDate      *time.Time         `json:"end_date,omitempty"`
	Status       string             `json:"status"`
}

func (c *ContractIn) ToContract() *Contract {
	return &Contract{
		TenantID:     c.TenantID,
		ProductID:    c.ProductID,
		ContractMode: c.ContractMode,
		StartDate:    c.StartDate,
		EndDate:      c.EndDate,
		Status:       c.Status,
	}
}

func (c *Contract) ToContractOut() *ContractOut {
	return &ContractOut{
		ID:           c.ID,
		TenantID:     c.TenantID,
		ProductID:    c.ProductID,
		ContractMode: c.ContractMode,
		StartDate:    c.StartDate,
		EndDate:      c.EndDate,
		Status:       c.Status,
	}
}
