package repository

import (
	"log"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/setting"
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
	db, err = sqlx.Connect("pgx", setting.AppSetting.DbConfig.Url)
	if err != nil {
		log.Fatalln("postgres.Setup err: %v", err)
		os.Exit(1)
	}
}

func Select(selectFields []string) string {
	if len(selectFields) == 0 {
		return "*"
	}

	// FIXME: Use strings.Join(selectFields, ", ") instead following
	var (
		s          = ""
		lastSelect = len(selectFields)
		count      int
		comma      = ", "
	)
	for _, field := range selectFields {
		count++
		if count == lastSelect {
			s += field
			break
		}
		s += field + comma
	}

	return s
}

func andWhere(criteria map[string][2]string, argsCount *int) (string, []interface{}) {
	if len(criteria) == 0 {
		return "", nil
	}
	var (
		conditions = " where "
		lastEl     = len(criteria)
		count      int
		and        = " and "
		args       []interface{}
		argNumb    string
	)
	for field, cond := range criteria {
		count++
		if count == lastEl {
			and = " "
		}
		argNumb = "$" + strconv.Itoa(*argsCount)
		conditions += field + cond[0] + argNumb + and
		args = append(args, cond[1])
		argNumb = ""
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
func queryBuilder(criteria map[string][2]string, order map[string]string, limit int, offset int) (string, []interface{}) {
	var (
		sql       string
		argsCount = 1
	)
	andWhere, args := andWhere(criteria, &argsCount)
	sql += andWhere + orderBy(order) + maxResult(limit) + offsetRows(offset)

	return sql, args
}

func IsComparisonOperator(o string) bool {
	return o == Equal || o == NotEqual || o == Great || o == GreatEqual || o == Less || o == LessEqual
}
