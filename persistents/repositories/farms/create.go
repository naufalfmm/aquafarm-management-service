package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) Create(ctx context.Context, farm dao.Farm) (dao.Farm, error) {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Create(&farm).
		Error(); err != nil {
		return dao.Farm{}, err
	}

	return farm, nil
}
