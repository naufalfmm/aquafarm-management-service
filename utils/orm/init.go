package orm

import (
	"fmt"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlConfig struct {
	Address  string
	Username string
	Password string
	DBName   string

	MaxIdleConnection int
	MaxOpenConnection int
	ConnMaxLifetime   time.Duration
}

func (c MySqlConfig) generateURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		c.Username, c.Password, c.Address, c.DBName)
}

func NewMysql(config MySqlConfig) (Orm, error) {
	db, err := gorm.Open(mysql.Open(config.generateURI()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(config.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(config.MaxOpenConnection)
	sqlDb.SetConnMaxLifetime(config.ConnMaxLifetime)

	o, err := NewOrm(db, consts.MySql)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
