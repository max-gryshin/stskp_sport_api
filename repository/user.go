package repository

import (
	"context"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/models"
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
