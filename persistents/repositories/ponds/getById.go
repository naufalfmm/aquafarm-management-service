package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetByID(ctx context.Context, id uint64) (dao.Pond, error) {
	var pond dao.Pond

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Preload("Farm").
		Where("id", id).
		Take(&pond).
		Error(); err != nil {
		return dao.Pond{}, err
	}

	return pond, nil
}
