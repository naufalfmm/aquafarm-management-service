package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) UpdateByID(ctx context.Context, id uint64, updateMap map[string]interface{}) error {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Where("id", id).
		Take(&dao.Farm{}).
		Updates(updateMap).
		Error(); err != nil {
		return err
	}

	return nil
}
