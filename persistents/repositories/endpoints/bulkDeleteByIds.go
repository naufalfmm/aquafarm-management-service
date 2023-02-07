package endpointsRepositories

import (
	"context"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) BulkDeleteByIDs(ctx context.Context, ids []uint64) error {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Model(&dao.Endpoint{}).
		Where("id IN ?", ids).
		UpdateColumns(map[string]interface{}{
			"deleted_at": time.Now(),
			"deleted_by": consts.SystemCreatedBy,
		}).
		Error(); err != nil {
		return err
	}

	return nil
}
