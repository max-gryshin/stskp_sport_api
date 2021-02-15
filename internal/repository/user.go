package repository

import (
	"fmt"
	"strings"

	"github.com/doug-martin/goqu/v9/exp"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // import the dialect

	"github.com/fatih/structs"
)

const tagName = "db"

func (repo *UserRepository) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	sql, _, err := repo.
		baseQuery.
		WithDialect("postgres").
		Where(exp.Ex{"user_name": username}).
		ToSQL()
	if err != nil {
		logging.Error(err)
		return user, err
	}
	err = repo.db.Get(&user, sql, username) // "select * from users where user_name = $1"
	if err != nil {
		logging.Error(err)
		return user, err
	}

	return user, nil
}

// get user by conditions
func FindUserBy(criteria map[string][2]string, order map[string]string, limit, offset int, selectFields []string) (models.Users, error) {
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
	err := db.Get(
		&user,
		strings.Join([]string{"select", Select(selectFields), "from \"user\" where id=$1"}, " "),
		id,
	)
	if err != nil {
		logging.Error(err)
		return user, err
	}

	return user, nil
}

// todo: do more flexible query
func UpdateUser(user *models.User) error {
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
