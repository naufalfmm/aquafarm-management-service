package pondsRepositories

import (
	"context"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

func (r repositories) DeleteByID(ctx context.Context, id uint64, loginDeleted token.Data) error {
	now := time.Now()
	if err := r.resources.MySql.Orm.
		Model(&dao.Pond{}).
		Where("id", id).
		UpdateColumns(map[string]interface{}{
			"deleted_at":   time.Now(),
			"deleted_by":   loginDeleted.CreatedBy(),
			"deleted_unix": now.Unix(),
		}).
		Error(); err != nil {
		return err
	}

	return nil
}
