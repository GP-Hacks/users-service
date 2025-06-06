package service_provider

import (
	app_cfg "github.com/GP-Hacks/users/internal/config"
	"github.com/GP-Hacks/users/internal/infrastructure/avatar_uploader"
	"github.com/GP-Hacks/users/internal/infrastructure/users_repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

func (s *ServiceProvider) AvatarUploader() *avatar_uploader.AvatarUploader {
	if s.avatarUploader == nil {
		log.Info().Msg("reg: " + app_cfg.Cfg.S3.Region + "\n")
		cfg := aws.Config{
			Region: app_cfg.Cfg.S3.Region,
			Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
				app_cfg.Cfg.S3.AccessKey,
				app_cfg.Cfg.S3.SecretKey,
				"",
			)),
			BaseEndpoint: &app_cfg.Cfg.S3.Endpoint,
		}
		cl := s3.NewFromConfig(cfg)

		s.avatarUploader = avatar_uploader.NewAvatarUploader(cl)
	}

	return s.avatarUploader
}

func (s *ServiceProvider) UsersRepository() *users_repository.UsersRepository {
	if s.usersRepository == nil {
		s.usersRepository = users_repository.NewUsersRepository(s.DB())
	}

	return s.usersRepository
}
