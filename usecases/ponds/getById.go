package pondsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (u usecases) GetByID(ctx context.Context, id uint64) (dao.Pond, error) {
	return u.persistents.Repositories.Ponds.GetByID(ctx, id)
}
