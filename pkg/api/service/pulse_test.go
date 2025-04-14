package service

import (
	"context"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service/mocks"
	"testing"
)

func TestPulse_Create(t *testing.T) {
	repo := mocks.NewMockPulseRepository()
	service := NewPulse(repo)

	pulseIn := &model.PulseIn{}
	_, err := service.Create(context.Background(), pulseIn)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestPulse_GetCurrentMonthConsumption(t *testing.T) {
	repo := mocks.NewMockPulseRepository()
	service := NewPulse(repo)

	_, err := service.GetCurrentMonthConsumption(context.Background(), "tenant1", "sku1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestPulse_GetAllResourcesConsumption(t *testing.T) {
	repo := mocks.NewMockPulseRepository()
	service := NewPulse(repo)

	_, err := service.GetAllResourcesConsumption(context.Background(), "tenant1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestPulse_AggregateAndStore(t *testing.T) {
	repo := mocks.NewMockPulseRepository()
	service := NewPulse(repo)

	err := service.AggregateAndStore(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
