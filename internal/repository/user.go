package repository

import (
	"context"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/models"
)

func FindUserByUsername(username string) (models.User, error) {
	row := db.QueryRow(context.Background(), "select * from \"user\" where user_name=$1", username)
	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.State, &user.CreatedAt, &user.Email)
	if err != nil {
		return *new(models.User), err
	}

	return user, nil
}

func CreateUser(user models.User) error {
	if _, err := db.Exec(
		context.Background(),
		"insert into \"user\"(user_name, password_hash, state, created_at, email) values ($1, $2, $3, $4, $5)",
		user.Username,
		user.Password,
		user.State,
		user.CreatedAt,
		user.Email,
	); err != nil {
		return err
	}

	return nil
}
