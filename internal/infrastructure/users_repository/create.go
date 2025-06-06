package users_repository

import (
	"context"

	"github.com/GP-Hacks/users/internal/models"
	"github.com/GP-Hacks/users/internal/services"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

func (r *UsersRepository) Create(ctx context.Context, u *models.User) error {
	q := `INSERT INTO users (id, email, first_name, last_name, surname, avatar_url, date_of_birth, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	if _, err := r.pool.Exec(ctx, q, u.ID, u.Email, u.FirstName, u.LastName, u.Surname, u.AvatarURL, u.DateOfBirth, u.Status, u.CreatedAt); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return services.AlreadyExists
			}
		}

		log.Error().Msg(err.Error())
		return services.InternalServerError
	}

	return nil
}
