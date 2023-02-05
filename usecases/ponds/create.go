package pondsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (u usecases) Create(ctx context.Context, req dto.CreatePondRequest) (dao.Pond, error) {
	farm, err := u.getFarm(ctx, req.FarmID, req.FarmCode)
	if err != nil {
		return dao.Pond{}, err
	}

	req.FarmID = farm.ID

	return u.persistents.Repositories.Ponds.Create(ctx, req.ToPond())
}

func (u usecases) getFarm(ctx context.Context, id uint64, code string) (dao.Farm, error) {
	if id != 0 {
		return u.persistents.Repositories.Farms.GetByID(ctx, id)
	}

	return u.persistents.Repositories.Farms.GetByCode(ctx, code)
}
