package endpointsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_GetAll(t *testing.T) {
	var (
		endpoints = dao.Endpoints{
			{
				ID: 1,
			},
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Endpoints
		mock.orm.EXPECT().Find(&data).DoAndReturn(func(data *dao.Endpoints, conds ...interface{}) interface{} {
			*data = endpoints
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetAll(mock.ctx)

		assert.Nil(t, err)
		assert.Equal(t, endpoints, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Endpoints
		mock.orm.EXPECT().Find(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when getting all endpoints",
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetAll(mock.ctx)

		assert.Equal(t, errAny, err)
		assert.Nil(t, res)
	})
}
