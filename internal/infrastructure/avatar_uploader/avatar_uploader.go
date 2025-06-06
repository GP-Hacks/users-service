package avatar_uploader

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AvatarUploader struct {
	client *s3.Client
}

func NewAvatarUploader(c *s3.Client) *AvatarUploader {
	return &AvatarUploader{
		client: c,
	}
}
