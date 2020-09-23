package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"log"
	"os"
	"strconv"
)

var db *sqlx.DB

const (
	ORDER_ASC  = "ASC"
	ORDER_DESC = "DESC"
)

func Setup() {
	var err error
	db, err = sqlx.Connect("pgx", setting.AppSetting.DbConfig.Url)
	if err != nil {
		log.Fatalln("postgres.Setup err: %v", err)
		os.Exit(1)
	}
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
		o         = ORDER_ASC
	)
	for field, order := range orderBy {
		count++
		if count == lastEl {
			comma = " "
		}
		if order == ORDER_DESC {
			o = ORDER_DESC
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

func queryBuilder(criteria map[string][2]string, order map[string]string, limit int, offset int) (string, []interface{}) {
	var (
		sql       string
		argsCount = 1
	)
	andWhere, args := andWhere(criteria, &argsCount)
	sql += andWhere + orderBy(order) + maxResult(limit) + offsetRows(offset)

	return sql, args
}
