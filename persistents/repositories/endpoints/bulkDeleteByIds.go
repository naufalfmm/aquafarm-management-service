package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) BulkDeleteByIDs(ctx context.Context, ids []uint64) error {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Model(&dao.Endpoint{}).
		Where("id IN ?", ids).
		UpdateColumns(map[string]interface{}{
			"deleted_at": timeNow(),
			"deleted_by": consts.SystemCreatedBy,
		}).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when bulk deleting by ids",
			zapLog.SetAttribute("ids", ids),
			zapLog.SetAttribute("error", err),
		)
		return err
	}

	return nil
}
