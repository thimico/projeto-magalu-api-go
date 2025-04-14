package controller

import (
	"context"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type HealthChecker interface {
	Check(ctx context.Context) error
}

type HealthHandler struct {
	hc HealthChecker
}

func NewHealthHandler(hc HealthChecker) *HealthHandler {
	return &HealthHandler{hc: hc}
}

// Healthcheck handles a health status check
// swagger:route GET /health health checkHealth
// Get health status
// responses:
//
//   200: description: Successful health check
//   500: description: Internal server error
func (p *HealthHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	error := p.hc.Check(ctx)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
