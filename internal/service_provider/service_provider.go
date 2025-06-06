package service_provider

import (
	"github.com/GP-Hacks/proto/pkg/api/auth"
	controllers "github.com/GP-Hacks/users/internal/controllers/grpc"
	"github.com/GP-Hacks/users/internal/infrastructure/auth_adapter"
	"github.com/GP-Hacks/users/internal/infrastructure/avatar_uploader"
	"github.com/GP-Hacks/users/internal/infrastructure/users_repository"
	"github.com/GP-Hacks/users/internal/services/user_service"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type ServiceProvider struct {
	authConnection *grpc.ClientConn
	db             *pgxpool.Pool
	authClient     auth.AuthServiceClient

	usersRepository *users_repository.UsersRepository
	avatarUploader  *avatar_uploader.AvatarUploader
	authAdapter     *auth_adapter.AuthAdapter

	usersService *user_service.UserService

	usersController *controllers.UserController
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
