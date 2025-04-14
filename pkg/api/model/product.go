package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"product_id"`
	Name       string             `bson:"name" json:"name"`
	Price      float64            `bson:"price" json:"price"`
	MetricUnit string             `bson:"metric_unit" json:"metric_unit"` // exemplo: "GB", "GB*hour", "req", "minute"
}

type ProductIn struct {
	Name       string  `json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
	MetricUnit string  `json:"metric_unit" validate:"required"`
}

type ProductOut struct {
	ID         primitive.ObjectID `json:"product_id"`
	Name       string             `json:"name"`
	Price      float64            `json:"price"`
	MetricUnit string             `json:"metric_unit"`
}

func (p *ProductIn) ToProduct() *Product {
	return &Product{
		Name:       p.Name,
		Price:      p.Price,
		MetricUnit: p.MetricUnit,
	}
}

func (p *Product) ToProductOut() *ProductOut {
	return &ProductOut{
		ID:         p.ID,
		Name:       p.Name,
		Price:      p.Price,
		MetricUnit: p.MetricUnit,
	}
}
