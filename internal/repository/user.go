package repository

import (
	"github.com/fatih/structs"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/models"
)

func FindUserByUsername(username string) (models.User, error) {
	user := models.User{}
	err := db.Get(&user, "select * from \"user\" where user_name=$1", username)
	if err != nil {
		logging.Error(err)
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) error {
	userMap := structs.Map(user)
	_, err := db.NamedExec(
		"INSERT INTO \"user\" (user_name, password_hash, state, created_at, email) VALUES (:Username,:Password,:State,:CreatedAt,:Email)",
		userMap,
	)
	if err != nil {
		return err
	}

	return nil
}

func UserAll() (models.Users, error) {
	users := models.Users{}
	err := db.Select(&users, "select * from \"user\"")
	if err != nil {
		logging.Error(err)
	}
	return users, err
}

func GetUserByID(id int) (models.User, error) {
	user := models.User{}
	err := db.Get(&user, "select * from \"user\" where id=$1", id)
	if err != nil {
		logging.Error(err)
		return user, err
	}

	return user, nil
}
