package pondsRepositories

import (
	"context"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

func (r repositories) DeleteByFarmID(ctx context.Context, farmID uint64, loginDeleted token.Data) error {
	now := time.Now()
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Model(&dao.Pond{}).
		Where("farm_id", farmID).
		UpdateColumns(map[string]interface{}{
			"deleted_at":   time.Now(),
			"deleted_by":   loginDeleted.CreatedBy(),
			"deleted_unix": now.Unix(),
		}).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when deleting ponds by farm id",
			zapLog.SetAttribute("farmID", farmID),
			zapLog.SetAttribute("error", err),
		)
		return err
	}

	return nil
}
