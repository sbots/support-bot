package service

import (
	"fmt"
	"support-bot/internal/infrastructure/errors"
	"support-bot/internal/models"
)

func (s *Service) UserSignUp(form *models.SignUpForm) (string, error) {
	tenant, err := s.repo.GetTenantByName(form.Company)
	if err != nil {
		return "", fmt.Errorf("getting tenant by name: %w", err)
	}

	yes, err := s.repo.TenantHasSuperuser(tenant.ID)
	if err != nil {
		return "", fmt.Errorf("checking for superuser failed: %w", err)
	}
	if yes {
		return "", fmt.Errorf("currently only one user per company is allowed")
	}

	user, err := models.NewUser(form.Name, form.Surname, form.Password, tenant.ID, form.Email, form.Phone)
	if err != nil {
		return "", fmt.Errorf("creating user failed: %w", err)
	}

	err = s.repo.UpsertUser(user)
	if err != nil {
		return "", fmt.Errorf("upserting user failed: %w", err)
	}
	token, err := s.auth.GetToken(user)
	if err != nil {
		return "", fmt.Errorf("getting authorisation token: %w", err)
	}
	return token, nil
}

func (s *Service) UserSignIn(form *models.SignInForm) (string, error) {
	user, err := s.repo.GetUserByEmail(form.Email)
	if err != nil {
		return "", err
	}

	if !user.ValidPassword(form.Password) {
		return "", errors.AccessDenied
	}

	token, err := s.auth.GetToken(user)
	if err != nil {
		return "", fmt.Errorf("getting authorisation token: %w", err)
	}
	return token, nil
}
