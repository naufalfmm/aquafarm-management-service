package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) UpdateByRequestID(ctx context.Context, requestID string, updateMap map[string]interface{}) error {
	if err := r.resources.MySql.GetDB().
		Model(&dao.EndpointLog{}).
		Where("request_id", requestID).
		Updates(updateMap).
		Error(); err != nil {
		return err
	}

	return nil
}
