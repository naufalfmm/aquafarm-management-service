package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type (
	Repositories interface {
		BulkCreate(ctx context.Context, data dao.Endpoints) (dao.Endpoints, error)
		GetAll(ctx context.Context) (dao.Endpoints, error)
		BulkDeleteByIDs(ctx context.Context, ids []uint64) error
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
