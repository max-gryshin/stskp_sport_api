package repository

import (
	"fmt"
	"strings"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
)

const tagName = "db"

// UserRepository is repository implementation for models.User
type UserRepository struct {
	BaseRepository
}

// NewUserRepository creates new instance of UserRepository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	table := "users"
	fields := utils.GetTagValue(models.User{}, tagName)

	return &UserRepository{BaseRepository{
		db:        db,
		table:     table,
		baseQuery: goqu.From(table).Select(fields).Prepared(true),
	}}
}

func (repo *UserRepository) GetByID(id int) (models.User, error) {
	user := models.User{}
	sql, _, err := repo.baseQuery.Where(exp.Ex{"id": id}).ToSQL()
	if err != nil {
		logging.Error(err)

		return user, err
	}

	err = db.Get(&user, sql, id)
	if err != nil {
		logging.Error(err)

		return user, err
	}

	return user, nil
}

func FindUserByUsername(username string) (models.User, error) {
	user := models.User{}
	err := db.Get(&user, "select * from \"user\" where user_name=$1", username)
	if err != nil {
		logging.Error(err)
		return user, err
	}

	return user, nil
}

func CreateUser(user *models.User) error {
	_, err := db.NamedExec(
		"INSERT INTO \"user\" (user_name, password_hash, state, created_at, email) VALUES (:Username,:Password,:State,:CreatedAt,:Email)",
		structs.Map(user), // FIXME: think about how to user struct instead map
	)
	if err != nil {
		return err
	}

	return nil
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
