package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type BaseRepository struct {
	db           *sqlx.DB
	table        string
	baseQuery    *goqu.SelectDataset
	queryBuilder goqu.DialectWrapper
}

func (repo *BaseRepository) execSelect(dest interface{}, data *goqu.SelectDataset) error {
	sql, params, err := data.ToSQL()
	if err != nil {
		return err
	}

	return db.Select(dest, sql, params...)
}

func (repo *BaseRepository) execInsert(data *goqu.InsertDataset) error {
	sql, params, err := data.ToSQL()
	if err != nil {
		return err
	}

	_ = repo.db.MustExec(sql, params...) // todo: solve the problem with duplicate keys

	return nil
}
