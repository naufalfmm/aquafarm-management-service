package endpointsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetByMethodPath(ctx context.Context, method, path string) (dao.Endpoint, error) {
	var endpoint dao.Endpoint
	if err := r.resources.MySql.GetDB().
		Where("method", method).
		Where("path", path).
		Take(&endpoint).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting endpoint by method and path",
			zapLog.SetAttribute("method", method),
			zapLog.SetAttribute("path", path),
			zapLog.SetAttribute("error", err),
		)
		return dao.Endpoint{}, err
	}

	return endpoint, nil
}
