package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
)

type Pulse interface {
	Create(ctx context.Context, pulse *model.PulseIn) (*model.Pulse, error)
	AggregateAndStore(ctx context.Context) error
	GetCurrentMonthConsumption(ctx context.Context, tenant string, sku string) (float64, error)
	GetAllResourcesConsumption(ctx context.Context, tenant string) (map[string]float64, error)
}

type PulseRepository interface {
	Create(ctx context.Context, pulse *model.Pulse) (*model.Pulse, error)
	AggregateAndStore(ctx context.Context) error
	GetCurrentMonthConsumption(ctx context.Context, tenant string, sku string) (float64, error)
	GetAllResourcesConsumption(ctx context.Context, tenant string) (map[string]float64, error)
}

type pulse struct {
	repository PulseRepository
}

func (s *pulse) GetCurrentMonthConsumption(ctx context.Context, tenant string, sku string) (float64, error) {
	return s.repository.GetCurrentMonthConsumption(ctx, tenant, sku)
}

func (s *pulse) GetAllResourcesConsumption(ctx context.Context, tenant string) (map[string]float64, error) {
	return s.repository.GetAllResourcesConsumption(ctx, tenant)
}

func NewPulse(repository PulseRepository) Pulse {
	return &pulse{repository: repository}
}

func (s *pulse) Create(ctx context.Context, in *model.PulseIn) (*model.Pulse, error) {
	pulseModel := in.ToPulse()
	pulse, err := s.repository.Create(ctx, pulseModel)
	if err != nil {
		return nil, err
	}
	return pulse, nil
}

func (s *pulse) AggregateAndStore(ctx context.Context) error {
	return s.repository.AggregateAndStore(ctx)
}
