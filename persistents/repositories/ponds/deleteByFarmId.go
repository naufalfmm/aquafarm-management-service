package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) DeleteByFarmID(ctx context.Context, farmID uint64, deletedBy string) error {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Model(&dao.Pond{}).
		Where("farm_id", farmID).
		UpdateColumns(map[string]interface{}{
			"deleted_at":   timeNow(),
			"deleted_by":   deletedBy,
			"deleted_unix": timeNow().Unix(),
		}).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when deleting ponds by farm id",
			zapLog.SetAttribute("farmID", farmID),
			zapLog.SetAttribute("deletedBy", deletedBy),
			zapLog.SetAttribute("error", err),
		)
		return err
	}

	return nil
}
