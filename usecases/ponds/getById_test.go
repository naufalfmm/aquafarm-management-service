package pondsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_GetByID(t *testing.T) {
	var (
		id   uint64   = 1
		pond dao.Pond = dao.Pond{
			ID: id,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.pondRepositories.EXPECT().GetByID(mock.ctx, id).Return(pond, nil)

		res, err := mock.usecases.GetByID(mock.ctx, id)

		assert.Nil(t, err)
		assert.Equal(t, pond, res)
	})
}
