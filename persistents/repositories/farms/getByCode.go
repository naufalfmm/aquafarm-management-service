package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetByCode(ctx context.Context, code string) (dao.Farm, error) {
	var farm dao.Farm

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Where("code", code).
		Take(&farm).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting farm by code",
			zapLog.SetAttribute("code", code),
			zapLog.SetAttribute("error", err),
		)
		return dao.Farm{}, err
	}

	return farm, nil
}
