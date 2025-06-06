package users_repository

import (
	"context"

	"github.com/GP-Hacks/users/internal/services"
	"github.com/rs/zerolog/log"
)

func (r *UsersRepository) UpdateAvatarURL(ctx context.Context, id int64, url string) error {
	q := `UPDATE users SET avatar_url = $1 WHERE id = $2`

	if _, err := r.pool.Exec(ctx, q, url, id); err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServerError
	}

	return nil
}
