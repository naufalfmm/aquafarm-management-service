package pondsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_GetAll(t *testing.T) {
	var (
		req   = dto.PondListRequest{}
		ponds = dao.Ponds{
			{
				ID: 1,
			},
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Ponds
		mock.orm.EXPECT().Preload("Farm").Return(mock.orm)
		mock.orm.EXPECT().Find(&data).DoAndReturn(func(data *dao.Ponds, conds ...interface{}) interface{} {
			*data = ponds
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetAll(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, ponds, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Ponds
		mock.orm.EXPECT().Preload("Farm").Return(mock.orm)
		mock.orm.EXPECT().Find(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when getting all ponds",
			zapLog.SetAttribute("req", req),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetAll(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Nil(t, res)
	})
}
