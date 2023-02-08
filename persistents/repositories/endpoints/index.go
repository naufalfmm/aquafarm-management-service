package endpointsRepositories

import (
	"context"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

var (
	timeNow = time.Now
)

type (
	Repositories interface {
		BulkCreate(ctx context.Context, data dao.Endpoints) (dao.Endpoints, error)
		GetAll(ctx context.Context) (dao.Endpoints, error)
		BulkDeleteByIDs(ctx context.Context, ids []uint64) error
		GetByMethodPath(ctx context.Context, method, path string) (dao.Endpoint, error)
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
