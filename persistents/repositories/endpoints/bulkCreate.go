package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) BulkCreate(ctx context.Context, data dao.Endpoints) (dao.Endpoints, error) {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Create(&data).
		Error(); err != nil {
		return nil, err
	}

	return data, nil
}
