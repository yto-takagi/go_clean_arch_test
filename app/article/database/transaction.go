package database

import (
	"context"
	"database/sql"
	"log"

	"go_clean_arch_test/app/article/transaction"

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
	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■ログテスト トランザクション1■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	tx := t.db.BeginTx(ctx, &sql.TxOptions{})
	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■ログテスト トランザクション2■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	ctx = context.WithValue(ctx, &txKey, tx)
	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■ログテスト トランザクション3■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	v, err := f(ctx)
	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■ログテスト トランザクション4■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	if err != nil {
		tx.Rollback()
		log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■ログテスト トランザクション5■■■■■■■■■■■■■■■■■■■■■■■■■■■■  " + err.Error())
		return nil, err
	}
	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■ログテスト トランザクション6■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
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
