package farmsRepositories

import (
	"context"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

var (
	timeNow = time.Now
)

type (
	Repositories interface {
		Create(ctx context.Context, farm dao.Farm) (dao.Farm, error)
		GetByID(ctx context.Context, id uint64) (dao.Farm, error)
		GetByCode(ctx context.Context, code string) (dao.Farm, error)
		DeleteByID(ctx context.Context, id uint64, deletedBy string) error
		GetAllPaginated(ctx context.Context, req dto.FarmPagingRequest) (dao.FarmsPagingResponse, error)
		UpdateByID(ctx context.Context, id uint64, updateMap map[string]interface{}) error
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
