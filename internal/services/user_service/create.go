package user_service

import (
	"context"
	"time"

	"github.com/GP-Hacks/users/internal/models"
)

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	user.AvatarURL = ""
	user.Status = models.DefaultUser
	user.CreatedAt = time.Now()
	if err := s.userRepository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
