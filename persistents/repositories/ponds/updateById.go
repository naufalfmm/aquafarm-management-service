package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) UpdateByID(ctx context.Context, id uint64, updateMap map[string]interface{}) error {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Where("id", id).
		Take(&dao.Pond{}).
		Updates(updateMap).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when updating pond by id",
			zapLog.SetAttribute("id", id),
			zapLog.SetAttribute("updateMap", updateMap),
			zapLog.SetAttribute("error", err),
		)
		return err
	}

	return nil
}
