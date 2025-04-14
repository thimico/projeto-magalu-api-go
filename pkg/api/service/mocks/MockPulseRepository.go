package mocks

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
)

type MockPulseRepository struct {
	pulses map[string]*model.Pulse
}

func NewMockPulseRepository() *MockPulseRepository {
	return &MockPulseRepository{
		pulses: make(map[string]*model.Pulse),
	}
}

func (m *MockPulseRepository) Create(ctx context.Context, pulse *model.Pulse) (*model.Pulse, error) {
	m.pulses[pulse.ID.String()] = pulse
	return pulse, nil
}

func (m *MockPulseRepository) AggregateAndStore(ctx context.Context) error {
	return nil
}

func (m *MockPulseRepository) GetCurrentMonthConsumption(ctx context.Context, tenant string, sku string) (float64, error) {
	return 0.0, nil
}

func (m *MockPulseRepository) GetAllResourcesConsumption(ctx context.Context, tenant string) (map[string]float64, error) {
	return map[string]float64{}, nil
}
