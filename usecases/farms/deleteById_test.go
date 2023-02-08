package farmsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_DeleteByID(t *testing.T) {
	var (
		id        uint64 = 1
		deletedBy string = "engineer.test@test.com"

		farm = dao.Farm{
			ID: id,
		}
	)

	t.Run("If no error, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, id).Return(farm, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)

		mock.farmRepositories.EXPECT().DeleteByID(mock.ctx, id, deletedBy).Return(nil)
		mock.pondRepositories.EXPECT().DeleteByFarmID(mock.ctx, id, deletedBy).Return(nil)

		mock.orm.EXPECT().Commit()

		err := mock.usecases.DeleteByID(mock.ctx, id, deletedBy)

		assert.Nil(t, err)
	})

	t.Run("If delete ponds by farm id error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, id).Return(farm, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)
		mock.orm.EXPECT().Rollback()

		mock.farmRepositories.EXPECT().DeleteByID(mock.ctx, id, deletedBy).Return(nil)
		mock.pondRepositories.EXPECT().DeleteByFarmID(mock.ctx, id, deletedBy).Return(errAny)

		err := mock.usecases.DeleteByID(mock.ctx, id, deletedBy)

		assert.Equal(t, errAny, err)
	})

	t.Run("If delete farm by id error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, id).Return(farm, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)
		mock.orm.EXPECT().Rollback()

		mock.farmRepositories.EXPECT().DeleteByID(mock.ctx, id, deletedBy).Return(errAny)

		err := mock.usecases.DeleteByID(mock.ctx, id, deletedBy)

		assert.Equal(t, errAny, err)
	})

	t.Run("If get farm by id error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, id).Return(dao.Farm{}, errAny)

		err := mock.usecases.DeleteByID(mock.ctx, id, deletedBy)

		assert.Equal(t, errAny, err)
	})
}
