package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
)

func (r repositories) GetAllPaginated(ctx context.Context, req dto.PondPagingRequest) (dao.PondsPagingResponse, error) {
	var (
		basePagingResp orm.BasePagingResponse
		ponds          dao.Ponds

		sortMap = map[string][]string{
			"code":        {"code"},
			"createdDate": {"created_at"},
		}
	)

	if err := req.Filter.Apply(r.resources.MySql.GetDB().WithContext(ctx)).
		Model(&dao.Pond{}).
		Paginate(ctx, orm.PaginateOptions{
			Paging:       req.PagingRequest,
			FieldSortMap: sortMap,
		}, &basePagingResp, &ponds).Error(); err != nil {
		return dao.PondsPagingResponse{}, err
	}

	return dao.PondsPagingResponse{
		BasePagingResponse: basePagingResp,
		Items:              ponds,
	}, nil
}
