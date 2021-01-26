package repository

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const (
	OrderAsc   = "ASC"
	OrderDesc  = "DESC"
	Equal      = "="
	NotEqual   = "!="
	Great      = ">"
	GreatEqual = ">="
	Less       = "<"
	LessEqual  = "<="
)

func Setup() {
	var err error
	db, err = sqlx.Connect("pgx", setting.AppSetting.DBConfig.URL)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
		os.Exit(1)
	}
}

func Select(selectFields []string) string {
	if len(selectFields) == 0 {
		return "*"
	}

	return strings.Join(selectFields, ", ")
}

func andWhere(criteria map[string][2]string, argsCount *int) (conditions string, args []interface{}) {
	lenCriteria := len(criteria)
	if lenCriteria == 0 {
		return "", nil
	}
	var (
		lastEl = lenCriteria
		count  int
		and    = " and "
	)
	conditions = " where "
	for field, cond := range criteria {
		count++
		if count == lastEl {
			and = " "
		}
		conditions += strings.Join([]string{field, cond[0], "$", strconv.Itoa(*argsCount), and}, "")
		args = append(args, cond[1])
		*argsCount++
	}

	return conditions, args
}

func orderBy(orderBy map[string]string) string {
	if len(orderBy) == 0 {
		return ""
	}
	var (
		orderCond = " order by "
		comma     = ", "
		lastEl    = len(orderBy)
		count     int
		o         = OrderAsc
	)
	for field, order := range orderBy {
		count++
		if count == lastEl {
			comma = " "
		}
		if order == OrderDesc {
			o = OrderDesc
		}
		orderCond += field + " " + o + comma
	}

	return orderCond
}

func offsetRows(offset int) string {
	return " offset " + strconv.Itoa(offset) + " rows "
}

func maxResult(maxResult int) string {
	if maxResult <= 0 {
		return ""
	}
	return " limit " + strconv.Itoa(maxResult)
}

// TODO:
// https://github.com/didi/gendry
// http://doug-martin.github.io/goqu/
// https://github.com/huandu/go-sqlbuilder
// https://github.com/Masterminds/squirrel
func queryBuilder(criteria map[string][2]string, order map[string]string, limit, offset int) (sql string, arguments []interface{}) {
	var argsCount = 1
	andWhere, args := andWhere(criteria, &argsCount)
	sql += andWhere + orderBy(order) + maxResult(limit) + offsetRows(offset)

	return sql, args
}

func IsComparisonOperator(o string) bool {
	return o == Equal || o == NotEqual || o == Great || o == GreatEqual || o == Less || o == LessEqual
}
