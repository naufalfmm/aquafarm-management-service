package endpointLogsRepositories

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_repository_GetAllReports(t *testing.T) {
	var (
		reports = dao.EndpointLogReports{
			{
				EndpointID: 1,
			},
		}

		gormDb = &gorm.DB{}
	)

	t.Run("If no error, it will return the list reports", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		var data dao.EndpointLogReports

		mock.orm.EXPECT().Table("endpoint_logs").Return(mock.orm)
		mock.orm.EXPECT().Select("endpoint_id",
			"user_agent",
			"ip_address",
			"SUM(end_at-start_at) as timelaps",
			"COUNT(*) as total").Return(mock.orm)
		mock.orm.EXPECT().Group("endpoint_id, user_agent, ip_address").Return(mock.orm)
		mock.orm.EXPECT().Gorm().Return(gormDb)

		mock.orm.EXPECT().Table("(?) as result", gormDb).Return(mock.orm)
		mock.orm.EXPECT().Preload("Endpoint").Return(mock.orm)
		mock.orm.EXPECT().Select(
			"endpoint_id",
			"SUM(total) as count",
			"COUNT(DISTINCT user_agent) as user_agent_distinct_count",
			"COUNT(DISTINCT ip_address) as ip_address_distinct_count",
			"SUM(timelaps)/SUM(total) as request_time_average",
		).Return(mock.orm)
		mock.orm.EXPECT().Group("endpoint_id").Return(mock.orm)
		mock.orm.EXPECT().Find(&data).DoAndReturn(func(data *dao.EndpointLogReports, conds ...interface{}) interface{} {
			*data = reports
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetAllReports(mock.ctx)

		assert.Nil(t, err)
		assert.Equal(t, reports, res)
	})

	t.Run("If orm return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.ctrl.Finish()

		mock.Before()

		var data dao.EndpointLogReports

		mock.orm.EXPECT().Table("endpoint_logs").Return(mock.orm)
		mock.orm.EXPECT().Select("endpoint_id",
			"user_agent",
			"ip_address",
			"SUM(end_at-start_at) as timelaps",
			"COUNT(*) as total").Return(mock.orm)
		mock.orm.EXPECT().Group("endpoint_id, user_agent, ip_address").Return(mock.orm)
		mock.orm.EXPECT().Gorm().Return(gormDb)

		mock.orm.EXPECT().Table("(?) as result", gormDb).Return(mock.orm)
		mock.orm.EXPECT().Preload("Endpoint").Return(mock.orm)
		mock.orm.EXPECT().Select(
			"endpoint_id",
			"SUM(total) as count",
			"COUNT(DISTINCT user_agent) as user_agent_distinct_count",
			"COUNT(DISTINCT ip_address) as ip_address_distinct_count",
			"SUM(timelaps)/SUM(total) as request_time_average",
		).Return(mock.orm)
		mock.orm.EXPECT().Group("endpoint_id").Return(mock.orm)
		mock.orm.EXPECT().Find(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)
		mock.logger.EXPECT().Error(mock.ctx,
			"error when getting all endpoint reports",
			zapLog.SetAttribute("error", errAny),
		)

		res, err := mock.repositories.GetAllReports(mock.ctx)

		assert.Equal(t, errAny, err)
		assert.Nil(t, res)
	})
}
