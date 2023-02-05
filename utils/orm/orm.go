package orm

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orm struct {
	g    *gorm.DB
	name string
}

func NewOrm(g *gorm.DB, name string) (orm, error) {
	return orm{
		g:    g,
		name: name,
	}, nil
}

func (o *orm) newImpl(g *gorm.DB) *orm {
	return &orm{
		g:    g,
		name: o.name,
	}
}

func (o *orm) Gorm() *gorm.DB {
	return o.g
}

func (o *orm) Error() error {
	return o.g.Error
}

func (o *orm) AddError(err error) error {
	return o.g.AddError(err)
}

func (o *orm) Assign(attrs ...interface{}) Orm {
	return o.newImpl(o.g.Assign(attrs...))
}

func (o *orm) Association(column string) *gorm.Association {
	return o.g.Association(column)
}

func (o *orm) Attrs(attrs ...interface{}) Orm {
	return o.newImpl(o.g.Attrs(attrs...))
}

func (o *orm) AutoMigrate(dst ...interface{}) error {
	return o.g.AutoMigrate()
}

func (o *orm) Begin(opts ...*sql.TxOptions) Orm {
	return o.newImpl(o.g.Begin(opts...))
}

func (o *orm) Clauses(conds ...clause.Expression) Orm {
	return o.newImpl(o.g.Clauses(conds...))
}

func (o *orm) Commit() Orm {
	return o.newImpl(o.g.Commit())
}

func (o *orm) Connection(fc func(tx Orm) error) error {
	return o.g.Connection(func(tx *gorm.DB) error {
		return fc(o)
	})
}

func (o *orm) Count(count *int64) Orm {
	return o.newImpl(o.g.Count(count))
}

func (o *orm) Create(value interface{}) Orm {
	return o.newImpl(o.g.Create(value))
}

func (o *orm) CreateInBatches(value interface{}, batchSize int) Orm {
	return o.newImpl(o.g.CreateInBatches(value, batchSize))
}

func (o *orm) DB() (*sql.DB, error) {
	return o.g.DB()
}

func (o *orm) Debug() Orm {
	return o.newImpl(o.g.Debug())
}

func (o *orm) Delete(value interface{}, conds ...interface{}) Orm {
	return o.newImpl(o.g.Delete(value, conds...))
}

func (o *orm) Distinct(args ...interface{}) Orm {
	return o.newImpl(o.g.Distinct(args...))
}

func (o *orm) Exec(sql string, values ...interface{}) Orm {
	return o.newImpl(o.g.Exec(sql, values...))
}

func (o *orm) Find(dest interface{}, conds ...interface{}) Orm {
	return o.newImpl(o.g.Find(dest, conds...))
}

func (o *orm) FindInBatches(dest interface{}, batchSize int, fc func(tx Orm, batch int) error) Orm {
	return o.newImpl(o.g.FindInBatches(dest, batchSize, func(tx *gorm.DB, batch int) error {
		return fc(o, batch)
	}))
}

func (o *orm) First(dest interface{}, conds ...interface{}) Orm {
	return o.newImpl(o.g.First(dest, conds...))
}

func (o *orm) FirstOrCreate(dest interface{}, conds ...interface{}) Orm {
	return o.newImpl(o.g.FirstOrCreate(dest, conds...))
}

func (o *orm) FirstOrInit(dest interface{}, conds ...interface{}) Orm {
	return o.newImpl(o.g.FirstOrInit(dest, conds...))
}

func (o *orm) Get(key string) (interface{}, bool) {
	return o.g.Get(key)
}

func (o *orm) Group(name string) Orm {
	return o.newImpl(o.g.Group(name))
}

func (o *orm) Having(query interface{}, args ...interface{}) Orm {
	return o.newImpl(o.g.Having(query, args...))
}

func (o *orm) InnerJoins(query string, args ...interface{}) Orm {
	return o.newImpl(o.g.InnerJoins(query, args...))
}

func (o *orm) InstanceGet(key string) (interface{}, bool) {
	return o.g.InstanceGet(key)
}

func (o *orm) InstanceSet(key string, value interface{}) Orm {
	return o.newImpl(o.g.InstanceSet(key, value))
}

func (o *orm) Joins(query string, args ...interface{}) Orm {
	return o.newImpl(o.g.Joins(query, args...))
}

func (o *orm) Last(dest interface{}, conds ...interface{}) Orm {
	return o.newImpl(o.g.Last(dest, conds...))
}

func (o *orm) Limit(limit int) Orm {
	return o.newImpl(o.g.Limit(limit))
}

func (o *orm) Migrator() gorm.Migrator {
	return o.g.Migrator()
}

func (o *orm) Model(value interface{}) Orm {
	return o.newImpl(o.g.Model(value))
}

func (o *orm) Not(query interface{}, args ...interface{}) Orm {
	return o.newImpl(o.g.Not(query, args...))
}

func (o *orm) Offset(offset int) Orm {
	return o.newImpl(o.g.Offset(offset))
}

func (o *orm) Omit(columns ...string) Orm {
	return o.newImpl(o.g.Omit(columns...))
}

func (o *orm) Or(query interface{}, args ...interface{}) Orm {
	return o.newImpl(o.g.Or(query, args...))
}

func (o *orm) Order(value interface{}) Orm {
	return o.newImpl(o.g.Order(value))
}

func (o *orm) Pluck(column string, dest interface{}) Orm {
	return o.newImpl(o.g.Pluck(column, dest))
}

func (o *orm) Preload(query string, args ...interface{}) Orm {
	return o.newImpl(o.g.Preload(query, args...))
}

func (o *orm) Raw(sql string, values ...interface{}) Orm {
	return o.newImpl(o.g.Raw(sql, values...))
}

func (o *orm) Rollback() Orm {
	return o.newImpl(o.g.Rollback())
}

func (o *orm) RollbackTo(name string) Orm {
	return o.newImpl(o.g.RollbackTo(name))
}

func (o *orm) Row() *sql.Row {
	return o.g.Row()
}

func (o *orm) Rows() (*sql.Rows, error) {
	return o.g.Rows()
}

func (o *orm) Save(value interface{}) Orm {
	return o.newImpl(o.g.Save(value))
}

func (o *orm) SavePoint(name string) Orm {
	return o.newImpl(o.g.SavePoint(name))
}

func (o *orm) Scan(dest interface{}) Orm {
	return o.newImpl(o.g.Scan(dest))
}

func (o *orm) ScanRows(rows *sql.Rows, dest interface{}) error {
	return o.g.ScanRows(rows, dest)
}

func (o *orm) Scopes(funcs ...func(tx Orm) Orm) Orm {
	newFuncs := make([]func(tx *gorm.DB) *gorm.DB, len(funcs))
	for i := 0; i < len(funcs); i++ {
		newFuncs[i] = func(tx *gorm.DB) *gorm.DB {
			return funcs[i](o).Gorm()
		}
	}

	return o.newImpl(o.g.Scopes(newFuncs...))
}

func (o *orm) Select(query interface{}, args ...interface{}) Orm {
	return o.newImpl(o.g.Select(query, args...))
}

func (o *orm) Session(config *gorm.Session) Orm {
	return o.newImpl(o.g.Session(config))
}

func (o *orm) Set(key string, value interface{}) Orm {
	return o.newImpl(o.g.Set(key, value))
}

func (o *orm) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return o.g.SetupJoinTable(model, field, joinTable)
}

func (o *orm) Table(name string, args ...interface{}) Orm {
	return o.newImpl(o.g.Table(name, args...))
}

func (o *orm) Take(dest interface{}, conds ...interface{}) Orm {
	return o.newImpl(o.g.Take(dest, conds...))
}

func (o *orm) ToSQL(queryFn func(tx Orm) Orm) string {
	return o.g.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return queryFn(o).Gorm()
	})
}

func (o *orm) Transaction(fc func(tx Orm) error, opts ...*sql.TxOptions) error {
	return o.g.Transaction(func(tx *gorm.DB) error {
		return fc(o)
	}, opts...)
}

func (o *orm) Unscoped() Orm {
	return o.newImpl(o.g.Unscoped())
}

func (o *orm) Update(column string, value interface{}) Orm {
	return o.newImpl(o.g.Update(column, value))
}

func (o *orm) UpdateColumn(column string, value interface{}) Orm {
	return o.newImpl(o.g.UpdateColumn(column, value))
}

func (o *orm) UpdateColumns(values interface{}) Orm {
	return o.newImpl(o.g.UpdateColumns(values))
}

func (o *orm) Updates(values interface{}) Orm {
	return o.newImpl(o.g.Updates(values))
}

func (o *orm) Use(plugin gorm.Plugin) error {
	return o.g.Use(plugin)
}

func (o *orm) Where(query interface{}, args ...interface{}) Orm {
	return o.newImpl(o.g.Where(query, args...))
}

func (o *orm) WithContext(ctx context.Context) Orm {
	return o.newImpl(o.g.WithContext(ctx))
}
