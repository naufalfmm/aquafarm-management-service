package farmsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_Create(t *testing.T) {
	var (
		farm = dao.Farm{
			Code: "FA-1",
		}
	)

	t.Run("If no error, it will return the created data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&farm).DoAndReturn(func(data *dao.Farm) interface{} {
			*data = farm
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.Create(mock.ctx, farm)

		assert.Nil(t, err)
		assert.Equal(t, farm, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&farm).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when creating farm",
			zapLog.SetAttribute("data", farm),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.Create(mock.ctx, farm)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Farm{}, res)
	})
}
