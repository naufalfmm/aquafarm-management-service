package pondsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (u usecases) Upsert(ctx context.Context, req dto.UpsertPondRequest) (dao.Pond, error) {
	// u.resources.MySql.SetPreloads(orm.SetPreload("Farm"))
	ponds, err := u.persistents.Repositories.Ponds.GetAll(ctx, dto.PondListRequest{
		ListPondFilterRequest: dto.ListPondFilterRequest{
			Code:     req.Code,
			FarmID:   req.FarmID,
			FarmCode: req.FarmCode,
		},
	})
	if err != nil {
		return dao.Pond{}, err
	}
	// u.resources.MySql.ResetPreloads()

	if len(ponds) == 0 {
		return u.Create(ctx, dto.CreatePondRequest{
			FarmID:      req.FarmID,
			FarmCode:    req.FarmCode,
			Code:        req.Code,
			Description: req.Description,
			Wide:        req.Wide,
			Long:        req.Long,
			Depth:       req.Depth,
			LoginData:   req.LoginData,
		})
	}

	req.FarmID = ponds[0].FarmID
	if err := u.persistents.Repositories.Ponds.UpdateByID(ctx, ponds[0].ID, req.ToUpdateMap()); err != nil {
		return dao.Pond{}, err
	}

	return u.persistents.Repositories.Ponds.GetByID(ctx, ponds[0].ID)
}
