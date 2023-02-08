package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetByCode(ctx context.Context, code string) (dao.Pond, error) {
	var pond dao.Pond

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Where("code", code).
		Take(&pond).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting pond by code",
			zapLog.SetAttribute("code", code),
			zapLog.SetAttribute("error", err),
		)
		return dao.Pond{}, err
	}

	return pond, nil
}
