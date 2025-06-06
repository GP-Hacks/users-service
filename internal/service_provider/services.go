package service_provider

import "github.com/GP-Hacks/users/internal/services/user_service"

func (s *ServiceProvider) UsersService() *user_service.UserService {
	if s.usersService == nil {
		s.usersService = user_service.NewUserService(s.AvatarUploader(), s.UsersRepository(), s.AuthAdapter())
	}

	return s.usersService
}
