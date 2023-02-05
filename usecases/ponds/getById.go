package pondsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func (u usecases) GetByID(ctx context.Context, id uint64) (dao.Pond, error) {
	u.resources.MySql.SetPreloads(orm.PreloadOpts{
		{
			Query: "Farm",
		},
	})

	return u.persistents.Repositories.Ponds.GetByID(ctx, id)
}
