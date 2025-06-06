package service_provider

import (
	controllers "github.com/GP-Hacks/users/internal/controllers/grpc"
)

func (s *ServiceProvider) UserController() *controllers.UserController {
	if s.usersController == nil {
		s.usersController = controllers.NewUserController(s.UsersService())
	}

	return s.usersController
}
