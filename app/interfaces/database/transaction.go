package database

import (
	"context"
	"database/sql"

	"go_clean_arch_test/app/transaction"

	"github.com/jinzhu/gorm"
)

var txKey = struct{}{}

type tx struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) transaction.Transaction {
	return &tx{db: db}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx := t.db.BeginTx(ctx, &sql.TxOptions{})
	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return v, nil
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}

func DoInTx(db *gorm.DB, f func(tx *gorm.DB) (interface{}, error)) (interface{}, error) {
	// start transaction
	tx := db.Begin()

	// execution
	v, err := f(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// commit
	tx.Commit()
	return v, nil
}
