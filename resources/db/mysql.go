package db

import (
	"fmt"

	"github.com/naufalfmm/aquafarm-management-service/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func generateURI(conf *config.EnvConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		conf.MySqlDbUsername, conf.MySqlDbPassword, conf.MySqlDbAddress, conf.MySqlDbName)
}

func NewMysql(config *config.EnvConfig) (*DB, error) {
	// db, err := gorm.Open(mysql.Open(generateURI(config)), &gorm.Config{})
	// if err != nil {
	// 	return nil, err
	// }

	// sqlDb, err := db.DB()
	// if err != nil {
	// 	return nil, err
	// }

	// sqlDb.SetMaxIdleConns(config.MySqlMaxIdleConnection)
	// sqlDb.SetMaxOpenConns(config.MySqlMaxOpenConnection)
	// sqlDb.SetConnMaxLifetime(config.MySqlConnMaxLifetime)

	// o, err := orm.NewOrm(db)
	// if err != nil {
	// 	return nil, err
	// }

	// return &DB{
	// 	Orm: &o,
	// }, nil

	o, err := orm.NewMysql(orm.MySqlConfig{
		Address:           config.MySqlDbAddress,
		Username:          config.MySqlDbUsername,
		Password:          config.MySqlDbPassword,
		DBName:            config.MySqlDbName,
		MaxIdleConnection: config.MySqlMaxIdleConnection,
		MaxOpenConnection: config.MySqlMaxOpenConnection,
		ConnMaxLifetime:   config.MySqlConnMaxLifetime,
	})
	if err != nil {
		return nil, err
	}

	return &DB{
		Orm: o,
	}, nil
}