package farmsRepositories

import (
	"context"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
	"github.com/stretchr/testify/assert"
)

func Test_repository_GetAllPaginated(t *testing.T) {
	var (
		req        = dto.FarmPagingRequest{}
		pagingResp = dao.FarmsPagingResponse{
			BasePagingResponse: orm.BasePagingResponse{
				CurrentPage: 1,
				TotalPage:   2,
				Count:       3,
				Limit:       2,
			},
			Items: dao.Farms{
				{ID: 1},
				{ID: 2},
			},
		}
	)

	t.Run("If no error, it will return the paginated farms", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var (
			basePagingResp orm.BasePagingResponse
			farms          dao.Farms
		)
		mock.orm.EXPECT().Model(&dao.Farm{}).Return(mock.orm)
		mock.orm.EXPECT().Paginate(mock.ctx, orm.PaginateOptions{
			Paging:       req.PagingRequest,
			FieldSortMap: sortMap,
		}, &basePagingResp, &farms).DoAndReturn(func(ctx context.Context, opt orm.PaginateOptions, basePagingResp *orm.BasePagingResponse, data *dao.Farms) interface{} {
			*basePagingResp = pagingResp.BasePagingResponse
			*data = pagingResp.Items

			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetAllPaginated(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, pagingResp, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var (
			basePagingResp orm.BasePagingResponse
			farms          dao.Farms
		)
		mock.orm.EXPECT().Model(&dao.Farm{}).Return(mock.orm)
		mock.orm.EXPECT().Paginate(mock.ctx, orm.PaginateOptions{
			Paging:       req.PagingRequest,
			FieldSortMap: sortMap,
		}, &basePagingResp, &farms).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when getting all paginated farms",
			zapLog.SetAttribute("req", req),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetAllPaginated(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.FarmsPagingResponse{}, res)
	})
}
