package repositories

import (
	endpointsRepositories "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/endpoints"
	farmsRepositories "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/farms"
	pondsRepositories "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/ponds"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type Repositories struct {
	Farms     farmsRepositories.Repositories
	Ponds     pondsRepositories.Repositories
	Endpoints endpointsRepositories.Repositories
}

func Init(res resources.Resources) (Repositories, error) {
	farmsRepo, err := farmsRepositories.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	pondsRepo, err := pondsRepositories.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	endpointsRepo, err := endpointsRepositories.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		Farms:     farmsRepo,
		Ponds:     pondsRepo,
		Endpoints: endpointsRepo,
	}, nil
}
