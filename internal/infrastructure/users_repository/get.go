package users_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/users/internal/models"
	"github.com/GP-Hacks/users/internal/services"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func (r *UsersRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	q := `SELECT id, email, first_name, last_name, surname, avatar_url, date_of_birth, status, created_at FROM users WHERE id = $1`

	var u models.User
	err := r.pool.QueryRow(ctx, q, id).Scan(
		&u.ID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.Surname,
		&u.AvatarURL,
		&u.DateOfBirth,
		&u.Status,
		&u.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}

	return &u, nil
}
