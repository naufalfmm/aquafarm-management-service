package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type (
	Repositories interface {
		Create(ctx context.Context, pond dao.Pond) (dao.Pond, error)
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
