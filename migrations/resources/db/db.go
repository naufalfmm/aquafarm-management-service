package db

import (
	"time"

	"github.com/naufalfmm/aquafarm-management-service/migrations/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func NewMysql(config *config.EnvConfig, log logger.Logger) (orm.Orm, error) {
	return orm.NewMysql([]orm.MysqlConfig{
		orm.WithAddress(config.MySqlDbAddress),
		orm.WithUsernamePassword(config.MySqlDbUsername, config.MySqlDbPassword),
		orm.WithDatabaseName(config.MySqlDbName),
		orm.WithLog(log, 200*time.Millisecond),
	}...)
}
