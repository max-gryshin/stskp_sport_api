package repository

import (
	"fmt"

	"github.com/max-gryshin/stskp_sport_api/internal/logging"
	"github.com/max-gryshin/stskp_sport_api/internal/models"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // import the dialect
)

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
