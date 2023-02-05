// package orm

// import (
// 	"context"
// 	"database/sql"

// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// type Association interface {
// 	Append(values ...interface{}) error
// 	Clear() error
// 	Count() (count int64)
// 	Delete(values ...interface{}) error
// 	Find(out interface{}, conds ...interface{}) error
// 	Replace(values ...interface{}) error
// }

// type processor interface {
// 	After(name string) callbacks
// 	Before(name string) callbacks
// 	Execute(db Orm) Orm
// 	Get(name string) func(Orm)
// 	Match(fc func(Orm) bool) callbacks
// 	Register(name string, fn func(Orm)) error
// 	Remove(name string) error
// 	Replace(name string, fn func(Orm)) error
// 	compile() (err error)
// }

// type callbacks interface {
// 	Create() processor
// 	Delete() processor
// 	Query() processor
// 	Raw() processor
// 	Row() processor
// 	Update() processor
// }

// type Orm interface {
// 	AddError(err error) error
// 	Assign(attrs ...interface{}) (tx Orm)
// 	Association(column string) Association
// 	Attrs(attrs ...interface{}) (tx Orm)
// 	AutoMigrate(dst ...interface{}) error
// 	Begin(opts ...*sql.TxOptions) Orm
// 	Callback() callbacks
// 	Clauses(conds ...clause.Expression) (tx Orm)
// 	Commit() Orm
// 	Connection(fc func(tx Orm) error) (err error)
// 	Count(count *int64) (tx Orm)
// 	Create(value interface{}) (tx Orm)
// 	CreateInBatches(value interface{}, batchSize int) (tx Orm)
// 	DB() (*sql.DB, error)
// 	Debug() (tx Orm)
// 	Delete(value interface{}, conds ...interface{}) (tx Orm)
// 	Distinct(args ...interface{}) (tx Orm)
// 	Exec(sql string, values ...interface{}) (tx Orm)
// 	Find(dest interface{}, conds ...interface{}) (tx Orm)
// 	FindInBatches(dest interface{}, batchSize int, fc func(tx Orm, batch int) error) Orm
// 	First(dest interface{}, conds ...interface{}) (tx Orm)
// 	FirstOrCreate(dest interface{}, conds ...interface{}) (tx Orm)
// 	FirstOrInit(dest interface{}, conds ...interface{}) (tx Orm)
// 	Get(key string) (interface{}, bool)
// 	Group(name string) (tx Orm)
// 	Having(query interface{}, args ...interface{}) (tx Orm)
// 	InnerJoins(query string, args ...interface{}) (tx Orm)
// 	InstanceGet(key string) (interface{}, bool)
// 	InstanceSet(key string, value interface{}) Orm
// 	Joins(query string, args ...interface{}) (tx Orm)
// 	Last(dest interface{}, conds ...interface{}) (tx Orm)
// 	Limit(limit int) (tx Orm)
// 	Migrator() gorm.Migrator
// 	Model(value interface{}) (tx Orm)
// 	Not(query interface{}, args ...interface{}) (tx Orm)
// 	Offset(offset int) (tx Orm)
// 	Omit(columns ...string) (tx Orm)
// 	Or(query interface{}, args ...interface{}) (tx Orm)
// 	Order(value interface{}) (tx Orm)
// 	Pluck(column string, dest interface{}) (tx Orm)
// 	Preload(query string, args ...interface{}) (tx Orm)
// 	Raw(sql string, values ...interface{}) (tx Orm)
// 	Rollback() Orm
// 	RollbackTo(name string) Orm
// 	Row() *sql.Row
// 	Rows() (*sql.Rows, error)
// 	Save(value interface{}) (tx Orm)
// 	SavePoint(name string) Orm
// 	Scan(dest interface{}) (tx Orm)
// 	ScanRows(rows *sql.Rows, dest interface{}) error
// 	Scopes(funcs ...func(Orm) Orm) (tx Orm)
// 	Select(query interface{}, args ...interface{}) (tx Orm)
// 	Session(config *gorm.Session) Orm
// 	Set(key string, value interface{}) Orm
// 	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
// 	Table(name string, args ...interface{}) (tx Orm)
// 	Take(dest interface{}, conds ...interface{}) (tx Orm)
// 	ToSQL(queryFn func(tx Orm) Orm) string
// 	Transaction(fc func(tx Orm) error, opts ...*sql.TxOptions) (err error)
// 	Unscoped() (tx Orm)
// 	Update(column string, value interface{}) (tx Orm)
// 	UpdateColumn(column string, value interface{}) (tx Orm)
// 	UpdateColumns(values interface{}) (tx Orm)
// 	Updates(values interface{}) (tx Orm)
// 	Use(plugin gorm.Plugin) error
// 	Where(query interface{}, args ...interface{}) (tx Orm)
// 	WithContext(ctx context.Context) Orm
// }

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
