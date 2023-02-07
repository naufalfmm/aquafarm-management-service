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
