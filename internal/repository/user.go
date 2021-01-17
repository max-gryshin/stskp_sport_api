package repository

import (
	"fmt"

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
	userMap := structs.Map(user) // FIXME: This line doesn't matter: https://jmoiron.github.io/sqlx/#namedParams
	_, err := db.NamedExec(
		"INSERT INTO \"user\" (user_name, password_hash, state, created_at, email) VALUES (:Username,:Password,:State,:CreatedAt,:Email)",
		userMap,
	)
	if err != nil {
		return err
	}

	return nil
}

// get user by conditions
func FindUserBy(criteria map[string][2]string, order map[string]string, limit int, offset int, selectFields []string) (models.Users, error) {
	var (
		sql   = "select " + Select(selectFields) + " from \"user\""
		users = models.Users{}
		err   error
	)
	query, args := queryBuilder(criteria, order, limit, offset)
	sql += query
	fmt.Printf(sql + "\n") // debug
	if err = db.Select(&users, sql, args...); err != nil {
		logging.Error(err)
	}
	return users, err
}

func GetUserByID(id int, selectFields []string) (models.User, error) {
	user := models.User{}
	err := db.Get(&user, "select "+Select(selectFields)+" from \"user\" where id=$1", id)
	if err != nil {
		logging.Error(err)
		return user, err
	}

	return user, nil
}

// todo: do more flexible query
func UpdateUser(user models.User) error {
	userMap := structs.Map(user)
	_, err := db.NamedExec(
		"UPDATE \"user\" SET user_name=:Username, state=:State, email=:Email WHERE id=:ID",
		userMap,
	)
	if err != nil {
		return err
	}

	return nil
}
