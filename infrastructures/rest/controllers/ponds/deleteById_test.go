package pondsControllers

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_controller_DeleteByID(t *testing.T) {
	var (
		id uint64 = 1
	)

	t.Run("If no error, it will return success", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", fmt.Sprintf("%d", id))

		mock.pondUsecases.EXPECT().DeleteByID(mock.ctx, id, mock.loginData.Email).Return(nil)

		err := mock.controllers.DeleteByID(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeSuccessResponse("Success", nil), mock.resRecorder.Body.String())
	})

	t.Run("If delete return record not found, it will return record not found", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", fmt.Sprintf("%d", id))

		expErr := gorm.ErrRecordNotFound

		mock.pondUsecases.EXPECT().DeleteByID(mock.ctx, id, mock.loginData.Email).Return(expErr)

		err := mock.controllers.DeleteByID(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(expErr.Error(), expErr), mock.resRecorder.Body.String())
	})

	t.Run("If delete return any error, it will return any error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", fmt.Sprintf("%d", id))

		expErr := errAny

		mock.pondUsecases.EXPECT().DeleteByID(mock.ctx, id, mock.loginData.Email).Return(expErr)

		err := mock.controllers.DeleteByID(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(expErr.Error(), expErr), mock.resRecorder.Body.String())
	})

	t.Run("If id missing, it will return id required error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", fmt.Sprintf("%d", 0))

		expErr := consts.ErrIdRequired

		err := mock.controllers.DeleteByID(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(expErr.Error(), expErr), mock.resRecorder.Body.String())
	})

	t.Run("If id is not number, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", "aaaa")

		expErr := errors.New("strconv.ParseUint: parsing \"aaaa\": invalid syntax")

		err := mock.controllers.DeleteByID(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(expErr.Error(), expErr), mock.resRecorder.Body.String())
	})
}
