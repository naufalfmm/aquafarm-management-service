package farmsControllers

import (
	"net/http"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
	"github.com/stretchr/testify/assert"
)

func Test_controller_GetAllPaginated(t *testing.T) {
	var (
		pagingResp = dao.FarmsPagingResponse{
			BasePagingResponse: orm.BasePagingResponse{
				Sorts: []string{"createdDate"},
			},
		}

		resp = dto.NewFarmPagingResponse(pagingResp)
	)

	t.Run("If no error, it will return the paginated data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.FarmPagingRequest{
			PagingRequest: orm.NewPagingRequest(mock.eCtx, []string{"createdDate"}),
		}

		mock.farmUsecases.EXPECT().GetAllPaginated(mock.ctx, req).Return(pagingResp, nil)

		err := mock.controllers.GetAllPaginated(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeSuccessResponse("Success", resp), mock.resRecorder.Body.String())
	})

	t.Run("If get all paginated return record not found, it will return record not found", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.FarmPagingRequest{
			PagingRequest: orm.NewPagingRequest(mock.eCtx, []string{"createdDate"}),
		}

		expErr := consts.ErrRecordNotFound

		mock.farmUsecases.EXPECT().GetAllPaginated(mock.ctx, req).Return(dao.FarmsPagingResponse{}, expErr)

		err := mock.controllers.GetAllPaginated(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(expErr.Error(), expErr), mock.resRecorder.Body.String())
	})

	t.Run("If get all paginated return any error, it will return any error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.FarmPagingRequest{
			PagingRequest: orm.NewPagingRequest(mock.eCtx, []string{"createdDate"}),
		}

		expErr := errAny

		mock.farmUsecases.EXPECT().GetAllPaginated(mock.ctx, req).Return(dao.FarmsPagingResponse{}, expErr)

		err := mock.controllers.GetAllPaginated(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(expErr.Error(), expErr), mock.resRecorder.Body.String())
	})
}
