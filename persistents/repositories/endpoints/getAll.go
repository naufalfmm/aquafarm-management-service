package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetAll(ctx context.Context) (dao.Endpoints, error) {
	var data dao.Endpoints

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Find(&data).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting all endpoints",
			zapLog.SetAttribute("error", err),
		)
		return nil, err
	}

	return data, nil
}
