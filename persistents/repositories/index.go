package repositories

import (
	endpointLogsRepositories "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/endpointLogs"
	endpointsRepositories "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/endpoints"
	farmsRepositories "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/farms"
	pondsRepositories "github.com/naufalfmm/aquafarm-management-service/persistents/repositories/ponds"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type Repositories struct {
	Farms        farmsRepositories.Repositories
	Ponds        pondsRepositories.Repositories
	Endpoints    endpointsRepositories.Repositories
	EndpointLogs endpointLogsRepositories.Repositories
}

func Init(res resources.Resources) (Repositories, error) {
	farmRepo, err := farmsRepositories.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	pondRepo, err := pondsRepositories.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	endpointRepo, err := endpointsRepositories.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	endpointLogRepo, err := endpointLogsRepositories.Init(res)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		Farms:        farmRepo,
		Ponds:        pondRepo,
		Endpoints:    endpointRepo,
		EndpointLogs: endpointLogRepo,
	}, nil
}
