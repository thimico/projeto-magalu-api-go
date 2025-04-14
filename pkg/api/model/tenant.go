package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tenant struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"tenant_id"`
	Name           string             `bson:"name" json:"name"`
	DocumentNumber string             `bson:"document_number" json:"document_number"`
}

type TenantIn struct {
	Name           string `json:"name" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required"`
}

type TenantOut struct {
	ID             primitive.ObjectID `json:"tenant_id"`
	Name           string             `json:"name"`
	DocumentNumber string             `json:"document_number"`
}

func (t *TenantIn) ToTenant() *Tenant {
	return &Tenant{
		Name:           t.Name,
		DocumentNumber: t.DocumentNumber,
	}
}

func (t *Tenant) ToTenantOut() *TenantOut {
	return &TenantOut{
		ID:             t.ID,
		Name:           t.Name,
		DocumentNumber: t.DocumentNumber,
	}
}
