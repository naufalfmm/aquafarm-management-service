package farmsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (u usecases) GetAllPaginated(ctx context.Context, req dto.FarmPagingRequest) (dao.FarmsPagingResponse, error) {
	resp, err := u.persistents.Repositories.Farms.GetAllPaginated(ctx, req)
	if err != nil {
		return dao.FarmsPagingResponse{}, err
	}

	if len(resp.Items) == 0 {
		return dao.FarmsPagingResponse{}, consts.ErrRecordNotFound
	}

	return resp, nil
}
