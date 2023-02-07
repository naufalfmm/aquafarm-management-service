package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetByRequestID(ctx context.Context, requestID string) (dao.EndpointLog, error) {
	var endpointLog dao.EndpointLog
	if err := r.resources.MySql.GetDB().
		Where("request_id", requestID).
		Take(&endpointLog).
		Error(); err != nil {
		return dao.EndpointLog{}, err
	}

	return endpointLog, nil
}
