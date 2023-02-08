package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetByRequestID(ctx context.Context, requestID string) (dao.EndpointLog, error) {
	var endpointLog dao.EndpointLog
	if err := r.resources.MySql.GetDB().
		Where("request_id", requestID).
		Take(&endpointLog).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting by request id",
			zapLog.SetAttribute("requestID", requestID),
			zapLog.SetAttribute("error", err),
		)
		return dao.EndpointLog{}, err
	}

	return endpointLog, nil
}
