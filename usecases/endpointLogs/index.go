package endpointLogsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/persistents"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type (
	Usecases interface {
		RecordRequestLog(ctx context.Context, req dto.RecordRequestLogRequest) (dao.EndpointLog, error)
		GetAllReports(ctx context.Context) (dao.EndpointLogReports, error)
	}

	usecases struct {
		persistents persistents.Persistents
		resources   resources.Resources
	}
)

func Init(persistents persistents.Persistents, resources resources.Resources) (Usecases, error) {
	return &usecases{
		persistents: persistents,
		resources:   resources,
	}, nil
}
