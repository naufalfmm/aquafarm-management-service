package pondsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_GetByID(t *testing.T) {
	var (
		id   uint64 = 1
		pond        = dao.Pond{
			ID: id,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Pond
		mock.orm.EXPECT().Preload("Farm").Return(mock.orm)
		mock.orm.EXPECT().Where("id", id).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).DoAndReturn(func(data *dao.Pond, conds ...interface{}) interface{} {
			*data = pond
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetByID(mock.ctx, id)

		assert.Nil(t, err)
		assert.Equal(t, pond, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Pond
		mock.orm.EXPECT().Preload("Farm").Return(mock.orm)
		mock.orm.EXPECT().Where("id", id).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when getting pond by id",
			zapLog.SetAttribute("id", id),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetByID(mock.ctx, id)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Pond{}, res)
	})
}
