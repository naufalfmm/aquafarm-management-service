package endpointsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (u usecases) BulkUpsertEndpoints(ctx context.Context, req dto.BulkUpsertEndpointsRequest) error {
	endpoints, err := u.persistents.Repositories.Endpoints.GetAll(ctx)
	if err != nil {
		return err
	}

	u.resources.MySql.StartTransaction(ctx)
	defer u.resources.MySql.RollbackTransaction(ctx)

	deletedIDs := []uint64{}
	for _, endpoint := range endpoints {
		foundReq := req.Endpoints.FindByMethodPath(endpoint.Method, endpoint.Path)
		if foundReq.Method != "" {
			continue
		}

		deletedIDs = append(deletedIDs, endpoint.ID)
	}

	if err := u.persistents.Repositories.Endpoints.BulkDeleteByIDs(ctx, deletedIDs); err != nil {
		return err
	}

	createdEndpoints := dao.Endpoints{}
	for _, endpointReq := range req.Endpoints {
		foundEndpoint := endpoints.FindByMethodPath(endpointReq.Method, endpointReq.Path)
		if foundEndpoint.ID != 0 {
			continue
		}

		createdEndpoints = append(createdEndpoints, endpointReq.ToEndpoint())
	}

	if len(createdEndpoints) > 0 {
		if _, err := u.persistents.Repositories.Endpoints.BulkCreate(ctx, createdEndpoints); err != nil {
			return err
		}
	}

	u.resources.MySql.CommitTransaction(ctx)

	return nil
}
