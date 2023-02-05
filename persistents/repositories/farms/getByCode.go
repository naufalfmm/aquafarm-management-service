package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetByCode(ctx context.Context, code string) (dao.Farm, error) {
	var farm dao.Farm

	if err := r.resources.MySql.GetDB().
		Where("code", code).
		Take(&farm).
		Error(); err != nil {
		return dao.Farm{}, err
	}

	return farm, nil
}
