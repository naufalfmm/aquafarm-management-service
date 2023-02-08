package endpointLogsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_Create(t *testing.T) {
	var (
		endpointLog = dao.EndpointLog{
			EndpointID: 1,
		}
	)

	t.Run("if create success, it will return the created data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&endpointLog).DoAndReturn(func(data *dao.EndpointLog) interface{} {
			*data = endpointLog
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.Create(mock.ctx, endpointLog)

		assert.Nil(t, err)
		assert.Equal(t, endpointLog, res)
	})

	t.Run("if orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&endpointLog).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when creating endpoint log",
			zapLog.SetAttribute("data", endpointLog),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.Create(mock.ctx, endpointLog)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.EndpointLog{}, res)
	})
}
