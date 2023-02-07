package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) Create(ctx context.Context, data dao.EndpointLog) (dao.EndpointLog, error) {
	if err := r.resources.MySql.GetDB().Create(&data).Error(); err != nil {
		return dao.EndpointLog{}, err
	}

	return data, nil
}
