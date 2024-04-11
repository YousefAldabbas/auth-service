package repository

import (
	"database/sql"

	"github.com/YousefAldabbas/auth-service/pkg/model"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	DB *sql.DB
}

func (r UserRepo) GetUserByUUID(userUUID string) (model.User, error) {
	q := `SELECT * FROM users WHERE uuid = $1`

	var user model.User

	err := r.DB.QueryRow(q, userUUID).Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil
}

func (r UserRepo) InsertUser(user *model.User) error {
	q := `INSERT INTO users (uuid, username, email, password) values ($1, $2, $3, $4) RETURNING id`

	err := r.DB.Ping()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(q, &user.UUID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepo) GetUserByUsername(username string) (model.User, error) {
	q := `Select * from users where username=$1 LIMIT 1`

	var user model.User
	err := r.DB.QueryRow(q, username).Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
