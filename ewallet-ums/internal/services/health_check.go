package services

import "ewallet-microservices/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepository
}

func (h *HealthCheck) HealthCheckService() (string, error) {
	return "Service healthy", nil
}
