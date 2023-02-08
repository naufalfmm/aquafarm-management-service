package farmsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func (r repositories) GetAllPaginated(ctx context.Context, req dto.FarmPagingRequest) (dao.FarmsPagingResponse, error) {
	var (
		basePagingResp orm.BasePagingResponse
		farms          dao.Farms

		sortMap = map[string][]string{
			"code":        {"code"},
			"createdDate": {"created_at"},
		}
	)

	if err := req.Filter.Apply(r.resources.MySql.GetDB().WithContext(ctx)).
		Model(&dao.Farm{}).
		Paginate(ctx, orm.PaginateOptions{
			Paging:       req.PagingRequest,
			FieldSortMap: sortMap,
		}, &basePagingResp, &farms).Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting all paginated farms",
			zapLog.SetAttribute("req", req),
			zapLog.SetAttribute("error", err),
		)
		return dao.FarmsPagingResponse{}, err
	}

	return dao.FarmsPagingResponse{
		BasePagingResponse: basePagingResp,
		Items:              farms,
	}, nil
}
