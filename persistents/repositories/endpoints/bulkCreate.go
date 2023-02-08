package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) BulkCreate(ctx context.Context, data dao.Endpoints) (dao.Endpoints, error) {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Create(&data).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when bulk creating endpoints",
			zapLog.SetAttribute("endpoints", data),
			zapLog.SetAttribute("error", err),
		)
		return nil, err
	}

	return data, nil
}
