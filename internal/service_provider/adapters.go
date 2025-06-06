package service_provider

import "github.com/GP-Hacks/users/internal/infrastructure/auth_adapter"

func (s *ServiceProvider) AuthAdapter() *auth_adapter.AuthAdapter {
	if s.authAdapter == nil {
		s.authAdapter = auth_adapter.NewAuthAdapter(s.AuthClient())
	}

	return s.authAdapter
}
