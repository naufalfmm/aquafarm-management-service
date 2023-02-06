package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type (
	Repositories interface {
		Create(ctx context.Context, pond dao.Pond) (dao.Pond, error)
		DeleteByID(ctx context.Context, id uint64, loginDeleted token.Data) error
		DeleteByFarmID(ctx context.Context, farmID uint64, loginDeleted token.Data) error
		GetByID(ctx context.Context, id uint64) (dao.Pond, error)
		GetAllPaginated(ctx context.Context, req dto.PondPagingRequest) (dao.PondsPagingResponse, error)
		GetByCode(ctx context.Context, code string) (dao.Pond, error) // DELETED
		UpdateByID(ctx context.Context, id uint64, updateMap map[string]interface{}) error
		GetAll(ctx context.Context, req dto.PondListRequest) (dao.Ponds, error)
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
