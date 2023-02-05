package orm

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Orm interface {
	Gorm() *gorm.DB
	Error() error
	SetPreloads(opts ...preloadOpt) Orm

	AddError(err error) error
	Assign(attrs ...interface{}) Orm
	Association(column string) *gorm.Association
	Attrs(attrs ...interface{}) Orm
	AutoMigrate(dst ...interface{}) error
	Begin(opts ...*sql.TxOptions) Orm
	Clauses(conds ...clause.Expression) Orm
	Commit() Orm
	Connection(fc func(tx Orm) error) error
	Count(count *int64) Orm
	Create(value interface{}) Orm
	CreateInBatches(value interface{}, batchSize int) Orm
	DB() (*sql.DB, error)
	Debug() Orm
	Delete(value interface{}, conds ...interface{}) Orm
	Distinct(args ...interface{}) Orm
	Exec(sql string, values ...interface{}) Orm
	Find(dest interface{}, conds ...interface{}) Orm
	FindInBatches(dest interface{}, batchSize int, fc func(tx Orm, batch int) error) Orm
	First(dest interface{}, conds ...interface{}) Orm
	FirstOrCreate(dest interface{}, conds ...interface{}) Orm
	FirstOrInit(dest interface{}, conds ...interface{}) Orm
	Get(key string) (interface{}, bool)
	Group(name string) Orm
	Having(query interface{}, args ...interface{}) Orm
	InnerJoins(query string, args ...interface{}) Orm
	InstanceGet(key string) (interface{}, bool)
	InstanceSet(key string, value interface{}) Orm
	Joins(query string, args ...interface{}) Orm
	Last(dest interface{}, conds ...interface{}) Orm
	Limit(limit int) Orm
	Migrator() gorm.Migrator
	Model(value interface{}) Orm
	Not(query interface{}, args ...interface{}) Orm
	Offset(offset int) Orm
	Omit(columns ...string) Orm
	Or(query interface{}, args ...interface{}) Orm
	Order(value interface{}) Orm
	Pluck(column string, dest interface{}) Orm
	Preload(query string, args ...interface{}) Orm
	Raw(sql string, values ...interface{}) Orm
	Rollback() Orm
	RollbackTo(name string) Orm
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Save(value interface{}) Orm
	SavePoint(name string) Orm
	Scan(dest interface{}) Orm
	ScanRows(rows *sql.Rows, dest interface{}) error
	Scopes(funcs ...func(tx Orm) Orm) Orm
	Select(query interface{}, args ...interface{}) Orm
	Session(config *gorm.Session) Orm
	Set(key string, value interface{}) Orm
	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
	Table(name string, args ...interface{}) Orm
	Take(dest interface{}, conds ...interface{}) Orm
	ToSQL(queryFn func(tx Orm) Orm) string
	Transaction(fc func(tx Orm) error, opts ...*sql.TxOptions) error
	Unscoped() Orm
	Update(column string, value interface{}) Orm
	UpdateColumn(column string, value interface{}) Orm
	UpdateColumns(values interface{}) Orm
	Updates(values interface{}) Orm
	Use(plugin gorm.Plugin) error
	Where(query interface{}, args ...interface{}) Orm
	WithContext(ctx context.Context) Orm
}
