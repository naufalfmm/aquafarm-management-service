package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) DeleteByID(ctx context.Context, id uint64, deletedBy string) error {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Model(&dao.Pond{}).
		Where("id", id).
		UpdateColumns(map[string]interface{}{
			"deleted_at":   timeNow(),
			"deleted_by":   deletedBy,
			"deleted_unix": timeNow().Unix(),
		}).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when deleting pond by id",
			zapLog.SetAttribute("id", id),
			zapLog.SetAttribute("deletedBy", deletedBy),
			zapLog.SetAttribute("error", err),
		)
		return err
	}

	return nil
}
