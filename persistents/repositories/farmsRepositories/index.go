package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type (
	Repositories interface {
		Create(ctx context.Context, farm dao.Farm) (dao.Farm, error)
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
