package farmsUsecases

import (
	"context"
	"errors"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"gorm.io/gorm"
)

func (u usecases) Upsert(ctx context.Context, req dto.UpsertFarmRequest) (dao.Farm, error) {
	farm, err := u.persistents.Repositories.Farms.GetByCode(ctx, req.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dao.Farm{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u.persistents.Repositories.Farms.Create(ctx, req.ToFarm())
	}

	if err := u.persistents.Repositories.Farms.UpdateByID(ctx, farm.ID, req.ToUpdateMap()); err != nil {
		return dao.Farm{}, err
	}

	return u.persistents.Repositories.Farms.GetByID(ctx, farm.ID)
}
