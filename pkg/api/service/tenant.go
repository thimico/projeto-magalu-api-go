package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
)

type Tenant interface {
	Create(ctx context.Context, tenant *model.TenantIn) (*model.TenantOut, error)
	GetByID(ctx context.Context, id string) (*model.TenantOut, error)
}

type TenantRepository interface {
	Create(ctx context.Context, tenant *model.Tenant) (*model.Tenant, error)
	GetByID(ctx context.Context, id string) (*model.Tenant, error)
}

type tenant struct {
	repository TenantRepository
}

func NewTenant(repository TenantRepository) Tenant {
	return &tenant{repository: repository}
}

func (s *tenant) Create(ctx context.Context, in *model.TenantIn) (*model.TenantOut, error) {
	tenantModel := in.ToTenant()
	tenant, err := s.repository.Create(ctx, tenantModel)
	if err != nil {
		return nil, err
	}
	return tenant.ToTenantOut(), nil
}

func (s *tenant) GetByID(ctx context.Context, id string) (*model.TenantOut, error) {
	tenant, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return tenant.ToTenantOut(), nil
}
