package endpointLogsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (u usecases) GetAllReports(ctx context.Context) (dao.EndpointLogReports, error) {
	return u.persistents.Repositories.EndpointLogs.GetAllReports(ctx)
}
