package endpointsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_BulkCreate(t *testing.T) {
	var (
		endpoints = dao.Endpoints{
			{
				ID: 1,
			},
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&endpoints).DoAndReturn(func(data *dao.Endpoints) interface{} {
			*data = endpoints
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.BulkCreate(mock.ctx, endpoints)

		assert.Nil(t, err)
		assert.Equal(t, endpoints, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&endpoints).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when bulk creating endpoints",
			zapLog.SetAttribute("endpoints", endpoints),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.BulkCreate(mock.ctx, endpoints)

		assert.Equal(t, errAny, err)
		assert.Nil(t, res)
	})
}
