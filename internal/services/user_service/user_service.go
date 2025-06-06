package user_service

import (
	"context"

	"github.com/GP-Hacks/users/internal/models"
)

type (
	IAuthAdapter interface {
		VerifyToken(ctx context.Context, token string) (int64, error)
	}

	IAvatarUploader interface {
		Upload(ctx context.Context, id int64, avatar []byte) (string, error)
	}

	IUserRepository interface {
		Create(ctx context.Context, user *models.User) error
		GetByID(ctx context.Context, id int64) (*models.User, error)
		UpdateAvatarURL(ctx context.Context, id int64, url string) error
		Update(ctx context.Context, usr *models.User) error
	}

	UserService struct {
		avatarUploader IAvatarUploader
		userRepository IUserRepository
		authAdapter    IAuthAdapter
	}
)

func NewUserService(au IAvatarUploader, ur IUserRepository, a IAuthAdapter) *UserService {
	return &UserService{
		avatarUploader: au,
		userRepository: ur,
		authAdapter:    a,
	}
}
