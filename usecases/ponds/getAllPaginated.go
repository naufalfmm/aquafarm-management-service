package pondsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func (u usecases) GetAllPaginated(ctx context.Context, req dto.PondPagingRequest) (dao.PondsPagingResponse, error) {
	u.resources.MySql.SetPreloads(orm.SetPreload("Farm"))

	resp, err := u.persistents.Repositories.Ponds.GetAllPaginated(ctx, req)
	if err != nil {
		return dao.PondsPagingResponse{}, err
	}

	if len(resp.Items) == 0 {
		return dao.PondsPagingResponse{}, consts.ErrRecordNotFound
	}

	return resp, nil
}
