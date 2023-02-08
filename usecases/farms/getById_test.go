package farmsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_GetByID(t *testing.T) {
	var (
		id   uint64   = 1
		farm dao.Farm = dao.Farm{
			ID: id,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, id).Return(farm, nil)

		res, err := mock.usecases.GetByID(mock.ctx, id)

		assert.Nil(t, err)
		assert.Equal(t, farm, res)
	})
}
