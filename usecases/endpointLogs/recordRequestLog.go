package endpointLogsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (u usecases) RecordRequestLog(ctx context.Context, req dto.RecordRequestLogRequest) (dao.EndpointLog, error) {
	endpoint, err := u.persistents.Repositories.Endpoints.GetByMethodPath(ctx, req.Method, req.Path)
	if err != nil {
		return dao.EndpointLog{}, err
	}

	req.EndpointID = endpoint.ID

	return u.persistents.Repositories.EndpointLogs.Create(ctx, req.ToEndpointLog())
}
