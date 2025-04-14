package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Pulse struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"pulse_id"`
	Tenant     string             `bson:"tenant" json:"tenant"`
	ProductSKU string             `bson:"product_sku" json:"product_sku"`
	UsedAmount float64            `bson:"used_amount" json:"used_amount"`
	UseUnity   string             `bson:"use_unity" json:"use_unity"`
	Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
}

type PulseIn struct {
	Tenant     string  `json:"tenant" validate:"required"`
	ProductSKU string  `json:"product_sku" validate:"required"`
	UsedAmount float64 `json:"used_amount" validate:"required"`
	UseUnity   string  `json:"use_unity" validate:"required"`
}

func (p *PulseIn) ToPulse() *Pulse {
	return &Pulse{
		Tenant:     p.Tenant,
		ProductSKU: p.ProductSKU,
		UsedAmount: p.UsedAmount,
		UseUnity:   p.UseUnity,
		Timestamp:  time.Now(),
	}
}