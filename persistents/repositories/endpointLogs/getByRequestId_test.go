package endpointLogsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_GetByRequestID(t *testing.T) {
	var (
		requestID string = "aaa-aaa-aaa"

		endpointLog = dao.EndpointLog{
			RequestID: requestID,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		var data dao.EndpointLog
		mock.orm.EXPECT().Where("request_id", requestID).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).DoAndReturn(func(data *dao.EndpointLog, conds ...interface{}) interface{} {
			*data = endpointLog
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetByRequestID(mock.ctx, requestID)

		assert.Nil(t, err)
		assert.Equal(t, endpointLog, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		var data dao.EndpointLog
		mock.orm.EXPECT().Where("request_id", requestID).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when getting by request id",
			zapLog.SetAttribute("requestID", requestID),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetByRequestID(mock.ctx, requestID)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.EndpointLog{}, res)
	})
}
