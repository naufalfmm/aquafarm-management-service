package farmsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_UpdateByID(t *testing.T) {
	var (
		id        uint64 = 1
		updateMap        = map[string]interface{}{
			"postal_code": "17158",
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Where("id", id).Return(mock.orm)
		mock.orm.EXPECT().Take(&dao.Farm{}).Return(mock.orm)
		mock.orm.EXPECT().Updates(updateMap).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(nil)

		err := mock.repositories.UpdateByID(mock.ctx, id, updateMap)

		assert.Nil(t, err)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		mock.orm.EXPECT().Where("id", id).Return(mock.orm)
		mock.orm.EXPECT().Take(&dao.Farm{}).Return(mock.orm)
		mock.orm.EXPECT().Updates(updateMap).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when updating farm by id",
			zapLog.SetAttribute("id", id),
			zapLog.SetAttribute("updateMap", updateMap),
			zapLog.SetAttribute("error", errAny),
		)

		err := mock.repositories.UpdateByID(mock.ctx, id, updateMap)

		assert.Equal(t, errAny, err)
	})
}
