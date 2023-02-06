package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetByID(ctx context.Context, id uint64) (dao.Farm, error) {
	var farm dao.Farm

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Where("id", id).
		Take(&farm).
		Error(); err != nil {
		return dao.Farm{}, err
	}

	return farm, nil
}
