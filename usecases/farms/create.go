package farmsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (u usecases) Create(ctx context.Context, req dto.CreateFarmRequest) (dao.Farm, error) {
	return u.persistents.Repositories.Farms.Create(ctx, req.ToFarm())
}
