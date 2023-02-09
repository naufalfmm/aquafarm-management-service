package farmsControllers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/stretchr/testify/assert"
)

func Test_controller_Upsert(t *testing.T) {
	var (
		bodyReq = dto.UpsertFarmRequest{
			Code: "A1",
		}

		farm = dao.Farm{
			ID:   1,
			Code: bodyReq.Code,
		}

		resp = dto.NewFarmResponse(farm)
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.UpsertFarmRequest{
			Code: "A1",

			LoginData: mock.loginData,
		}

		mock.SetRequestBody(bodyReq)

		mock.validator.EXPECT().Validate(&req).Return(nil)
		mock.farmUsecases.EXPECT().Upsert(mock.ctx, req).Return(farm, nil)

		err := mock.controllers.Upsert(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeSuccessResponse("Success", resp), mock.resRecorder.Body.String())
	})

	t.Run("If upsert error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.UpsertFarmRequest{
			Code: "A1",

			LoginData: mock.loginData,
		}

		mock.SetRequestBody(bodyReq)

		mock.validator.EXPECT().Validate(&req).Return(nil)
		mock.farmUsecases.EXPECT().Upsert(mock.ctx, req).Return(dao.Farm{}, errAny)

		err := mock.controllers.Upsert(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(errAny.Error(), errAny), mock.resRecorder.Body.String())
	})

	t.Run("If validate error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.UpsertFarmRequest{
			Code: "A1",

			LoginData: mock.loginData,
		}

		mock.SetRequestBody(bodyReq)

		mock.validator.EXPECT().Validate(&req).Return(errAny)

		err := mock.controllers.Upsert(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(errAny.Error(), errAny), mock.resRecorder.Body.String())
	})

	t.Run("If bind error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		expErr := errors.New("code=415, message=Unsupported Media Type")

		mock.SetRequestBody(bodyReq)
		mock.eCtx.Request().Header.Del(consts.ContentTypeHeader)

		err := mock.controllers.Upsert(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(expErr.Error(), expErr), mock.resRecorder.Body.String())
	})
}
