package avatar_uploader

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/GP-Hacks/users/internal/config"
	"github.com/GP-Hacks/users/internal/services"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

func (a *AvatarUploader) Upload(ctx context.Context, id int64, avatar []byte) (string, error) {
	key := fmt.Sprintf("avatars/%d_%d", id, time.Now().Unix())

	contentType := detectContentType(avatar)

	_, err := a.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(config.Cfg.S3.Bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        bytes.NewReader(avatar),
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return "", services.InternalServerError
	}

	return config.Cfg.S3.Endpoint + "/" + config.Cfg.S3.Bucket + "/" + key, nil
}

func detectContentType(data []byte) string {
	if len(data) >= 512 {
		return http.DetectContentType(data[:512])
	}
	return http.DetectContentType(data)
}
