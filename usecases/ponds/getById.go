package pondsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (u usecases) GetByID(ctx context.Context, id uint64) (dao.Pond, error) {
	// u.resources.MySql.SetPreloads(orm.SetPreload("Farm"))

	return u.persistents.Repositories.Ponds.GetByID(ctx, id)
}
