package endpointsRepositories

import (
	"net/http"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_repository_GetByMethodPath(t *testing.T) {
	var (
		method = http.MethodGet
		path   = "/v1/ponds"

		endpoint = dao.Endpoint{
			Method: method,
			Path:   path,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Endpoint
		mock.orm.EXPECT().Where("method", method).Return(mock.orm)
		mock.orm.EXPECT().Where("path", path).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).DoAndReturn(func(data *dao.Endpoint, conds ...interface{}) interface{} {
			*data = endpoint
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetByMethodPath(mock.ctx, method, path)

		assert.Nil(t, err)
		assert.Equal(t, endpoint, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.Before()

		var data dao.Endpoint
		mock.orm.EXPECT().Where("method", method).Return(mock.orm)
		mock.orm.EXPECT().Where("path", path).Return(mock.orm)
		mock.orm.EXPECT().Take(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx, "error when getting endpoint by method and path",
			zapLog.SetAttribute("method", method),
			zapLog.SetAttribute("path", path),
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetByMethodPath(mock.ctx, method, path)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Endpoint{}, res)
	})
}
