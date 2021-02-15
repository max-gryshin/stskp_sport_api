package repository

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
)

const maxItemsPerPage = 100

// UserRepository is repository implementation for models.User
type UserRepository struct {
	BaseRepository
}

// NewUserRepository creates new instance of UserRepository
func NewUserRepository(db *sqlx.DB, queryBuilder goqu.DialectWrapper) *UserRepository {
	table := `users`
	fields := utils.GetTagValue(models.User{}, tagName)
	baseQuery := queryBuilder.From(table).Select(fields...).Prepared(true)

	return &UserRepository{BaseRepository{
		db:           db,
		table:        table,
		baseQuery:    baseQuery,
		queryBuilder: queryBuilder,
	}}
}

func (repo *UserRepository) GetByID(id int) (models.User, error) {
	user := models.User{}
	sql, params, err := repo.baseQuery.Where(exp.Ex{"id": id}).ToSQL()
	if err != nil {
		logging.Error(err)

		return user, err
	}

	err = db.Get(&user, sql, params...)
	if err != nil {
		logging.Error(err)

		return user, err
	}

	return user, nil
}

func (repo *UserRepository) GetUsers() (models.Users, error) {
	var users = models.Users{}
	query := repo.baseQuery.Limit(maxItemsPerPage)
	sql, p, err := query.ToSQL()
	if err != nil {
		return users, err
	}

	err = repo.db.Select(&users, sql, p...)

	return users, err
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	query := repo.
		baseQuery.
		Insert().
		Into("users").
		Cols("user_name", "password_hash", "state", "created_at").
		Vals(goqu.Vals{user.Username, user.Password, user.State, user.CreatedAt})

	return repo.execInsert(query)
}
