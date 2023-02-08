package orm

import (
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(configs ...MysqlConfig) (Orm, error) {
	conf, err := defaultMysqlConfig()
	if err != nil {
		return nil, err
	}

	for _, config := range configs {
		config.Apply(&conf)
	}

	db, err := gorm.Open(mysql.Open(conf.generateURI()), conf.ToGormConfig())
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(conf.maxIdleConnection)
	sqlDb.SetMaxOpenConns(conf.maxOpenConnection)
	sqlDb.SetConnMaxLifetime(conf.connMaxLifetime)

	o, err := NewOrm(db, consts.MySql)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
