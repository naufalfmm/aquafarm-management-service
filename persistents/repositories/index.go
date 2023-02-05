package repositories

import (
	farmsRepository "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/farmsRepositories"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type Repositories struct {
	Farms farmsRepository.Repositories
}

func Init(res resources.Resources) (Repositories, error) {
	farmsRepo, err := farmsRepository.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		Farms: farmsRepo,
	}, nil
}
