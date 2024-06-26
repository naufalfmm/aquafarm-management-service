package db

import (
	"fmt"

	"github.com/naufalfmm/aquafarm-management-service/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func NewMysql(config *config.EnvConfig, log logger.Logger) (*DB, error) {
	confs := []orm.MysqlConfig{
		orm.WithAddress(fmt.Sprintf("%s:%s", config.MySqlDbHost, config.MySqlDbPort)),
		orm.WithUsernamePassword(config.MySqlDbUsername, config.MySqlDbPassword),
		orm.WithDatabaseName(config.MySqlDbName),
		orm.WithMaxIdleConnection(config.MySqlMaxIdleConnection),
		orm.WithMaxOpenConnection(config.MySqlMaxOpenConnection),
		orm.WithConnMaxLifetime(config.MySqlConnMaxLifetime),
		orm.WithRetry(config.MySqlRetry, config.MySqlWaitSleep),
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
