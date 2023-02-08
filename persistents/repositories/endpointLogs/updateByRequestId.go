package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) UpdateByRequestID(ctx context.Context, requestID string, updateMap map[string]interface{}) error {
	if err := r.resources.MySql.GetDB().
		Model(&dao.EndpointLog{}).
		Where("request_id", requestID).
		Updates(updateMap).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when updating by request id",
			zapLog.SetAttribute("requestID", requestID),
			zapLog.SetAttribute("updateMap", updateMap),
			zapLog.SetAttribute("error", err),
		)
		return err
	}

	return nil
}
