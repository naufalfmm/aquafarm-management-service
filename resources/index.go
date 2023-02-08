package resources

import (
	"github.com/naufalfmm/aquafarm-management-service/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/resources/db"
	jwToken "github.com/naufalfmm/aquafarm-management-service/resources/jwt"
	"github.com/naufalfmm/aquafarm-management-service/resources/log"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
	"github.com/naufalfmm/aquafarm-management-service/utils/validator"
)

type Resources struct {
	Config    *config.EnvConfig
	Logger    logger.Logger
	MySql     *db.DB
	Validator validator.Validator
	JWT       jwt.JWT
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

	validator, err := validator.NewV9()
	if err != nil {
		return Resources{}, err
	}

	jwtImp, err := jwToken.NewJWT(conf)
	if err != nil {
		return Resources{}, err
	}

	return Resources{
		Config:    conf,
		Logger:    logger,
		MySql:     mysql,
		Validator: validator,
		JWT:       jwtImp,
	}, nil
}
