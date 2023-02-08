package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) Create(ctx context.Context, farm dao.Farm) (dao.Farm, error) {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Create(&farm).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when creating farm",
			zapLog.SetAttribute("data", farm),
			zapLog.SetAttribute("error", err),
		)
		return dao.Farm{}, err
	}

	return farm, nil
}
