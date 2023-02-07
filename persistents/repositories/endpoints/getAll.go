package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetAll(ctx context.Context) (dao.Endpoints, error) {
	var data dao.Endpoints

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Find(&data).
		Error(); err != nil {
		return nil, err
	}

	return data, nil
}
