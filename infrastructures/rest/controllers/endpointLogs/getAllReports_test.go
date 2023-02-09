package endpointLogsControllers

import (
	"net/http"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/stretchr/testify/assert"
)

func Test_controller_GetAllReports(t *testing.T) {
	var (
		reports = dao.EndpointLogReports{
			{
				EndpointID: 1,
				Endpoint: dao.Endpoint{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
			},
		}

		resp = dto.NewEndpointLogReportResponseMap(reports)
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.endpointLogUsecases.EXPECT().GetAllReports(mock.ctx).Return(reports, nil)

		err := mock.controllers.GetAllReports(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeSuccessResponse("Success", resp), mock.resRecorder.Body.String())
	})

	t.Run("If get all reports error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.endpointLogUsecases.EXPECT().GetAllReports(mock.ctx).Return(nil, errAny)

		err := mock.controllers.GetAllReports(mock.eCtx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(errAny.Error(), errAny), mock.resRecorder.Body.String())
	})
}
