package farmsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func (u usecases) GetByID(ctx context.Context, id uint64) (dao.Farm, error) {
	u.resources.MySql.SetPreloads(orm.SetPreload("Ponds"))

	return u.persistents.Repositories.Farms.GetByID(ctx, id)
}
