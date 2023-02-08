package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) Create(ctx context.Context, pond dao.Pond) (dao.Pond, error) {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Create(&pond).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when creating pond",
			zapLog.SetAttribute("data", pond),
			zapLog.SetAttribute("error", err),
		)
		return dao.Pond{}, err
	}

	return pond, nil
}
