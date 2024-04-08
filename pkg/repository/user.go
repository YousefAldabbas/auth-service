package repository

import (
	"context"

	"github.com/YousefAldabbas/auth-service/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func (ur UserRepo) GetUserByUUID(userUUID string) (model.User, error) {
	q := `SELECT * FROM users WHERE uuid = $1`

	var user model.User

	err := ur.DB.QueryRow(context.Background(), q, userUUID).Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil
}

func (ur UserRepo) InsertUser(user *model.User) error {
    q := `INSERT INTO users (uuid, username, email, password) values ($1, $2, $3, $4) RETURNING id`


	err := ur.DB.Ping(context.Background())
	if err != nil {
		return err
	}



	_, err = ur.DB.Exec(context.Background(), q, &user.UUID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return err
	}

    return nil
}