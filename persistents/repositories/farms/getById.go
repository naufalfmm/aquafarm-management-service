package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetByID(ctx context.Context, id uint64) (dao.Farm, error) {
	var farm dao.Farm

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Preload("Ponds").
		Where("id", id).
		Take(&farm).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting farm by id",
			zapLog.SetAttribute("id", id),
			zapLog.SetAttribute("error", err),
		)
		return dao.Farm{}, err
	}

	return farm, nil
}
