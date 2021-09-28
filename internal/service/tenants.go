package service

import (
	"fmt"
	"support-bot/internal/models"
)

func (s *Service) NewTenant(name string) (*models.Tenant, error) {
	tenant := models.NewTenant(name)
	err := s.repo.UpsertTenant(tenant)
	if err != nil {
		return nil, fmt.Errorf("upserting tenant: %s", err.Error())
	}
	return tenant, nil
}
