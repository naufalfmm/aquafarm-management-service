package db

import "github.com/naufalfmm/aquafarm-management-service/utils/orm"

type Transaction struct {
	trx    orm.Orm
	isUsed bool
}

func (trx *Transaction) Ignore() {
	trx.isUsed = false
}

func (trx *Transaction) Acknowledge() {
	trx.isUsed = true
}

func (trx *Transaction) Begin(orm orm.Orm) {
	trx.trx = orm.Begin()
	trx.isUsed = true
}

func (trx *Transaction) Commit() {
	if !trx.isUsed {
		return
	}

	trx.trx.Commit()
	trx.isUsed = false
}

func (trx *Transaction) Rollback() {
	if !trx.isUsed {
		return
	}

	trx.trx.Rollback()
	trx.isUsed = false
}
