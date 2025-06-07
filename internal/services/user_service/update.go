package user_service

import (
	"context"

	"github.com/GP-Hacks/users/internal/models"
)

func (s *UserService) Update(ctx context.Context, token string, upd *models.User) error {
	id, err := s.authAdapter.VerifyToken(ctx, token)
	if err != nil {
		return err
	}

	upd.ID = id
	if err := s.userRepository.Update(ctx, upd); err != nil {
		return err
	}

	return nil
}
