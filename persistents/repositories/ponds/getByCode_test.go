package pondsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_GetByCode(t *testing.T) {
	var (
		code = "A1"
		pond = dao.Pond{
			ID:   1,
			Code: code,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Pond
		mock.orm.EXPECT().Where("code", code).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).DoAndReturn(func(data *dao.Pond, conds ...interface{}) interface{} {
			*data = pond
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetByCode(mock.ctx, code)

		assert.Nil(t, err)
		assert.Equal(t, pond, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Pond
		mock.orm.EXPECT().Where("code", code).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when getting pond by code",
			zapLog.SetAttribute("code", code),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetByCode(mock.ctx, code)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Pond{}, res)
	})
}
