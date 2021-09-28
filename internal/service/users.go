package service

import (
	"context"
	"fmt"
	"support-bot/internal/infrastructure/errors"
	"support-bot/internal/models"
)

func (s *Service) GetUserInformation(ctx context.Context) (*models.User, error) {
	token := s.auth.GetServiceTokenFromContext(ctx)
	if token == nil {
		return nil, errors.AccessDenied
	}

	user, err := s.repo.GetUserByID(token.GetUserID())
	if err != nil {
		return nil, fmt.Errorf("getting user by id: %w", err)
	}
	return user, err
}
