package pondsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_DeleteByID(t *testing.T) {
	var (
		id        uint64 = 1
		deletedBy string = "engineer.test@test.com"

		pond = dao.Pond{
			ID: id,
		}
	)

	t.Run("If no error, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.pondRepositories.EXPECT().GetByID(mock.ctx, id).Return(pond, nil)
		mock.pondRepositories.EXPECT().DeleteByID(mock.ctx, id, deletedBy).Return(nil)

		err := mock.usecases.DeleteByID(mock.ctx, id, deletedBy)

		assert.Nil(t, err)
	})

	t.Run("If get farm by id error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.pondRepositories.EXPECT().GetByID(mock.ctx, id).Return(dao.Pond{}, errAny)

		err := mock.usecases.DeleteByID(mock.ctx, id, deletedBy)

		assert.Equal(t, errAny, err)
	})
}
