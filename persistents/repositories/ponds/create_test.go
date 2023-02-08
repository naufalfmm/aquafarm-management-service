package pondsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_Create(t *testing.T) {
	var (
		pond = dao.Pond{
			Code: "A1",
		}
	)

	t.Run("If no error, it will return the created data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&pond).DoAndReturn(func(data *dao.Pond) interface{} {
			*data = pond
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.Create(mock.ctx, pond)

		assert.Nil(t, err)
		assert.Equal(t, pond, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Create(&pond).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when creating pond",
			zapLog.SetAttribute("data", pond),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.Create(mock.ctx, pond)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Pond{}, res)
	})
}
