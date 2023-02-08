package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

var (
	sortMap = map[string][]string{
		"code":        {"code"},
		"createdDate": {"created_at"},
	}
)

func (r repositories) GetAllPaginated(ctx context.Context, req dto.PondPagingRequest) (dao.PondsPagingResponse, error) {
	var (
		basePagingResp orm.BasePagingResponse
		ponds          dao.Ponds
	)

	if err := req.Filter.Apply(r.resources.MySql.GetDB().WithContext(ctx)).
		Model(&dao.Pond{}).
		Preload("Farm").
		Paginate(ctx, orm.PaginateOptions{
			Paging:       req.PagingRequest,
			FieldSortMap: sortMap,
		}, &basePagingResp, &ponds).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting all paginated ponds",
			zapLog.SetAttribute("req", req),
			zapLog.SetAttribute("error", err),
		)
		return dao.PondsPagingResponse{}, err
	}

	return dao.PondsPagingResponse{
		BasePagingResponse: basePagingResp,
		Items:              ponds,
	}, nil
}
