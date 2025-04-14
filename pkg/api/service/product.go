package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
)

type Product interface {
	Create(ctx context.Context, product *model.ProductIn) (*model.ProductOut, error)
	GetByID(ctx context.Context, id string) (*model.ProductOut, error)
}

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
}

type product struct {
	repository ProductRepository
}

func NewProduct(repository ProductRepository) Product {
	return &product{repository: repository}
}

func (s *product) Create(ctx context.Context, in *model.ProductIn) (*model.ProductOut, error) {
	productModel := in.ToProduct()
	product, err := s.repository.Create(ctx, productModel)
	if err != nil {
		return nil, err
	}
	return product.ToProductOut(), nil
}

func (s *product) GetByID(ctx context.Context, id string) (*model.ProductOut, error) {
	product, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product.ToProductOut(), nil
}
