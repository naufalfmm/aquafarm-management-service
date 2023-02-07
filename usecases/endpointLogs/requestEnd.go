package endpointLogsUsecases

import (
	"context"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (u usecases) RequestEnd(ctx context.Context, requestID string) (dao.EndpointLog, error) {
	if err := u.persistents.Repositories.EndpointLogs.UpdateByRequestID(ctx, requestID, map[string]interface{}{
		"end_at":     time.Now().UnixMilli(),
		"updated_at": time.Now(),
		"updated_by": consts.SystemCreatedBy,
	}); err != nil {
		return dao.EndpointLog{}, err
	}

	return u.persistents.Repositories.EndpointLogs.GetByRequestID(ctx, requestID)
}
