package users_repository

import (
	"context"

	"github.com/GP-Hacks/users/internal/models"
	"github.com/GP-Hacks/users/internal/services"
	"github.com/rs/zerolog/log"
)

func (r *UsersRepository) Update(ctx context.Context, usr *models.User) error {
	q := `UPDATE users SET first_name = $1, last_name = $2, surname = $3, date_of_birth = $4 WHERE id = $5`

	if _, err := r.pool.Exec(ctx, q, usr.FirstName, usr.LastName, usr.Surname, usr.DateOfBirth, usr.ID); err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServerError
	}

	return nil
}
