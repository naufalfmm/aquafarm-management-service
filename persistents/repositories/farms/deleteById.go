package farmsRepositories

import (
	"context"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

func (r repositories) DeleteByID(ctx context.Context, id uint64, loginDeleted token.Data) error {
	now := time.Now()
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Model(&dao.Farm{}).
		Where("id", id).
		UpdateColumns(map[string]interface{}{
			"deleted_at":   now,
			"deleted_by":   loginDeleted.CreatedBy(),
			"deleted_unix": now.Unix(),
		}).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when deleting farm by id",
			zapLog.SetAttribute("id", id),
			zapLog.SetAttribute("error", err),
		)
		return err
	}

	return nil
}
