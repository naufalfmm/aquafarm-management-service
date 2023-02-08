package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) Create(ctx context.Context, data dao.EndpointLog) (dao.EndpointLog, error) {
	if err := r.resources.MySql.GetDB().
		Create(&data).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when creating endpoint log",
			zapLog.SetAttribute("data", data),
			zapLog.SetAttribute("error", err),
		)
		return dao.EndpointLog{}, err
	}

	return data, nil
}
