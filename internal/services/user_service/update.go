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

	usr, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if upd.FirstName != "" {
		usr.FirstName = upd.FirstName
	}
	if upd.LastName != "" {
		usr.LastName = upd.LastName
	}
	if upd.Surname != "" {
		usr.Surname = upd.Surname
	}
	if !upd.DateOfBirth.IsZero() {
		usr.DateOfBirth = upd.DateOfBirth
	}

	if err := s.userRepository.Update(ctx, usr); err != nil {
		return err
	}

	return nil
}
