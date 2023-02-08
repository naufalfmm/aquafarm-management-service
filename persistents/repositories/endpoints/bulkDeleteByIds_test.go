package endpointsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_BulkDeleteByIDs(t *testing.T) {
	var (
		ids = []uint64{1, 2}
	)

	t.Run("If no error, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Model(&dao.Endpoint{}).Return(mock.orm)
		mock.orm.EXPECT().Where("id IN ?", ids).Return(mock.orm)
		mock.orm.EXPECT().UpdateColumns(map[string]interface{}{
			"deleted_at": mock.now,
			"deleted_by": consts.SystemCreatedBy,
		}).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(nil)

		err := mock.repositories.BulkDeleteByIDs(mock.ctx, ids)

		assert.Nil(t, err)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Model(&dao.Endpoint{}).Return(mock.orm)
		mock.orm.EXPECT().Where("id IN ?", ids).Return(mock.orm)
		mock.orm.EXPECT().UpdateColumns(map[string]interface{}{
			"deleted_at": mock.now,
			"deleted_by": consts.SystemCreatedBy,
		}).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when bulk deleting by ids",
			zapLog.SetAttribute("ids", ids),
			zapLog.SetAttribute("error", errAny),
		)

		err := mock.repositories.BulkDeleteByIDs(mock.ctx, ids)

		assert.Equal(t, errAny, err)
	})
}
