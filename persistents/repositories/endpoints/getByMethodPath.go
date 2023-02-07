package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetByMethodPath(ctx context.Context, method, path string) (dao.Endpoint, error) {
	var endpoint dao.Endpoint
	if err := r.resources.MySql.GetDB().
		Where("method", method).
		Where("path", path).
		Take(&endpoint).
		Error(); err != nil {
		return dao.Endpoint{}, err
	}

	return endpoint, nil
}
