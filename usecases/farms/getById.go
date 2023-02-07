package farmsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (u usecases) GetByID(ctx context.Context, id uint64) (dao.Farm, error) {
	return u.persistents.Repositories.Farms.GetByID(ctx, id)
}
