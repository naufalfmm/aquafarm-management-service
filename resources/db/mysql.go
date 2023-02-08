package db

import (
	"github.com/naufalfmm/aquafarm-management-service/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func NewMysql(config *config.EnvConfig, log logger.Logger) (*DB, error) {
	confs := []orm.MysqlConfig{
		orm.WithAddress(config.MySqlDbAddress),
		orm.WithUsernamePassword(config.MySqlDbUsername, config.MySqlDbPassword),
		orm.WithDatabaseName(config.MySqlDbName),
		orm.WithMaxIdleConnection(config.MySqlMaxIdleConnection),
		orm.WithMaxOpenConnection(config.MySqlMaxOpenConnection),
		orm.WithConnMaxLifetime(config.MySqlConnMaxLifetime),
	}

	if config.MySqlLogMode {
		confs = append(confs, orm.WithLog(log, config.MySqlLogSlowThreshold))
	}

	o, err := orm.NewMysql(confs...)
	if err != nil {
		return nil, err
	}

	return &DB{
		Orm: o,
	}, nil
}
