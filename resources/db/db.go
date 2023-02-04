package db

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

type DB struct {
	orm.Orm
	transaction *Transaction
}

func (db *DB) StartTransaction(ctx context.Context) {
	if db.transaction == nil {
		db.transaction = &Transaction{}
	}

	db.transaction.Begin(db.Orm)
}

func (db *DB) CommitTransaction(ctx context.Context) {
	if db.transaction == nil {
		return
	}

	db.transaction.Commit()
	db.transaction = nil
}

func (db *DB) RollbackTransaction(ctx context.Context) {
	if db.transaction == nil {
		return
	}

	db.transaction.Rollback()
	db.transaction = nil
}

func (db *DB) IgnoreTransaction() {
	if db.transaction == nil {
		return
	}

	db.transaction.Ignore()
}

func (db *DB) GetDB() orm.Orm {
	if db.transaction == nil {
		return db.Orm
	}

	if !db.transaction.isUsed {
		return db.Orm
	}

	return db.transaction.trx
}
