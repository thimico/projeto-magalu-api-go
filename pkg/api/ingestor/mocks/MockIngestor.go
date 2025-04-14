package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"projeto-magalu-api-go/pkg/api/model"
)

type Pulse struct {
	mock.Mock
}

func (m *Pulse) Create(ctx context.Context, pulse *model.PulseIn) (*model.Pulse, error) {
	args := m.Called(ctx, pulse)
	return args.Get(0).(*model.Pulse), args.Error(1)
}

func (m *Pulse) GetCurrentMonthConsumption(ctx context.Context, tenant string, sku string) (float64, error) {
	args := m.Called(ctx, tenant, sku)
	return args.Get(0).(float64), args.Error(1)
}

func (m *Pulse) AggregateAndStore(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *Pulse) GetAllResourcesConsumption(ctx context.Context, tenant string) (map[string]float64, error) {
	args := m.Called(ctx, tenant)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]float64), args.Error(1)
}
