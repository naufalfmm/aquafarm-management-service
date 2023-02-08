package farmsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_DeleteByID(t *testing.T) {
	var (
		id uint64 = 1

		deletedBy = "engineer.test@test.com"
	)

	t.Run("If no error, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Model(&dao.Farm{}).Return(mock.orm)
		mock.orm.EXPECT().Where("id", id).Return(mock.orm)
		mock.orm.EXPECT().UpdateColumns(map[string]interface{}{
			"deleted_at":   mock.now,
			"deleted_by":   deletedBy,
			"deleted_unix": mock.now.Unix(),
		}).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(nil)

		err := mock.repositories.DeleteByID(mock.ctx, id, deletedBy)

		assert.Nil(t, err)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Model(&dao.Farm{}).Return(mock.orm)
		mock.orm.EXPECT().Where("id", id).Return(mock.orm)
		mock.orm.EXPECT().UpdateColumns(map[string]interface{}{
			"deleted_at":   mock.now,
			"deleted_by":   deletedBy,
			"deleted_unix": mock.now.Unix(),
		}).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when deleting farm by id",
			zapLog.SetAttribute("id", id),
			zapLog.SetAttribute("deletedBy", deletedBy),
			zapLog.SetAttribute("error", errAny),
		)

		err := mock.repositories.DeleteByID(mock.ctx, id, deletedBy)

		assert.Equal(t, errAny, err)
	})
}
