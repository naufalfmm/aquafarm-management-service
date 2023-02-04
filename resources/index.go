package resources

import (
	"github.com/naufalfmm/aquafarm-management-service/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
	"github.com/naufalfmm/aquafarm-management-service/utils/validator"
)

type Resources struct {
	Config    *config.EnvConfig
	MySql     orm.Orm
	Validator validator.Validator
}

func Init() (Resources, error) {
	return Resources{}, nil
}
