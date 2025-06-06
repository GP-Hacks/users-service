package user_service

import (
	"context"

	"github.com/GP-Hacks/users/internal/models"
)

func (s *UserService) GetMe(ctx context.Context, token string) (*models.User, error) {
	id, err := s.authAdapter.VerifyToken(ctx, token)
	if err != nil {
		return nil, err
	}

	u, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
