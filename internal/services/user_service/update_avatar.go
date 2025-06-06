package user_service

import (
	"context"
)

func (s *UserService) UpdateAvatar(ctx context.Context, token string, avatar []byte) (string, error) {
	id, err := s.authAdapter.VerifyToken(ctx, token)
	if err != nil {
		return "", err
	}

	url, err := s.avatarUploader.Upload(ctx, id, avatar)
	if err != nil {
		return "", err
	}

	if err := s.userRepository.UpdateAvatarURL(ctx, id, url); err != nil {
		return "", err
	}

	return url, nil
}
