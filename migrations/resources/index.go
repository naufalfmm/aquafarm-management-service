package resources

import (
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources/db"
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources/log"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

type Resources struct {
	Config *config.EnvConfig
	Logger logger.Logger
	MySql  orm.Orm
}

func Init() (Resources, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return Resources{}, err
	}

	logger, err := log.NewLog()
	if err != nil {
		return Resources{}, err
	}

	mysql, err := db.NewMysql(conf, logger)
	if err != nil {
		return Resources{}, err
	}

	return Resources{
		Config: conf,
		MySql:  mysql,
	}, nil
}
