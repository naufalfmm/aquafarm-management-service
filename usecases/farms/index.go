package farmsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/persistents"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type (
	Usecases interface {
		Create(ctx context.Context, req dto.CreateFarmRequest) (dao.Farm, error)
		DeleteByID(ctx context.Context, id uint64, deletedBy string) error
		GetByID(ctx context.Context, id uint64) (dao.Farm, error)
		GetAllPaginated(ctx context.Context, req dto.FarmPagingRequest) (dao.FarmsPagingResponse, error)
		Upsert(ctx context.Context, req dto.UpsertFarmRequest) (dao.Farm, error)
	}

	usecases struct {
		persistents persistents.Persistents
		resources   resources.Resources
	}
)

func Init(persistents persistents.Persistents, resources resources.Resources) (Usecases, error) {
	return &usecases{
		persistents: persistents,
		resources:   resources,
	}, nil
}
