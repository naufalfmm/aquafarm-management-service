package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type (
	Repositories interface {
		Create(ctx context.Context, data dao.EndpointLog) (dao.EndpointLog, error)
		UpdateByRequestID(ctx context.Context, requestID string, updateMap map[string]interface{}) error
		GetByRequestID(ctx context.Context, requestID string) (dao.EndpointLog, error)
	}

	repositories struct {
		resources resources.Resources
	}
)

func Init(resources resources.Resources) (Repositories, error) {
	return &repositories{
		resources: resources,
	}, nil
}