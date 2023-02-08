package endpointLogsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_usecases_GetAllReports(t *testing.T) {
	var (
		reports = dao.EndpointLogReports{}
	)

	t.Run("If no error, it will return the data returned by the repository", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.endpointLogRepositories.EXPECT().GetAllReports(mock.ctx).Return(reports, nil)

		res, err := mock.usecases.GetAllReports(mock.ctx)

		assert.Nil(t, err)
		assert.Equal(t, reports, res)
	})
}
