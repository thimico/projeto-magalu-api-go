package mocks

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"projeto-magalu-api-go/pkg/api/model"
)

type PulseService struct {
	mock.Mock
}

func (_m *PulseService) AggregateAndStore(ctx context.Context) error {
	return nil
}

func (_m *PulseService) GetCurrentMonthConsumption(ctx context.Context, tenant string, sku string) (float64, error) {
	if tenant == "tenant1" && sku == "sku1" {
		return 100.5, nil
	}
	return 0, fmt.Errorf("invalid tenant or SKU")
}

func (_m *PulseService) GetAllResourcesConsumption(ctx context.Context, tenant string) (map[string]float64, error) {
	if tenant == "tenant1" {
		return map[string]float64{
			"resource1": 200.0,
			"resource2": 150.5,
		}, nil
	}
	return nil, fmt.Errorf("tenant not found")
}

func (_m *PulseService) Create(ctx context.Context, in *model.PulseIn) (*model.Pulse, error) {
	ret := _m.Called(ctx, in)
	var r0 *model.Pulse
	if rf, ok := ret.Get(0).(func(context.Context, *model.PulseIn) *model.Pulse); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(*model.Pulse)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.PulseIn) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}
