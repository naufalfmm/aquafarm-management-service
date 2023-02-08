package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetByID(ctx context.Context, id uint64) (dao.Pond, error) {
	var pond dao.Pond

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Preload("Farm").
		Where("id", id).
		Take(&pond).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when get by id", zapLog.SetAttribute("id", id), zapLog.SetAttribute("error", err))
		return dao.Pond{}, err
	}

	return pond, nil
}
