package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type (
	Repositories interface {
		Create(ctx context.Context, farm dao.Farm) (dao.Farm, error)
		GetByID(ctx context.Context, id uint64) (dao.Farm, error)
		GetByCode(ctx context.Context, code string) (dao.Farm, error)
		DeleteByID(ctx context.Context, id uint64, loginDeleted token.Data) error
		GetAllPaginated(ctx context.Context, req dto.FarmPagingRequest) (dao.FarmsPagingResponse, error)
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
