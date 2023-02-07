package endpointLogsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetAllReports(ctx context.Context) (dao.EndpointLogReports, error) {
	var (
		data dao.EndpointLogReports
		orm  = r.resources.MySql.GetDB()
	)

	reportOrm := orm.
		Table("endpoint_logs").
		Select("endpoint_id",
			"user_agent",
			"ip_address",
			"SUM(end_at-start_at) as timelaps",
			"COUNT(*) as total",
		).
		Group("endpoint_id, user_agent, ip_address")

	if err := orm.WithContext(ctx).
		Table("(?) as result", reportOrm.Gorm()).
		Preload("Endpoint").
		Select("endpoint_id",
			"SUM(total) as count",
			"COUNT(DISTINCT user_agent) as user_agent_distinct_count",
			"COUNT(DISTINCT ip_address) as ip_address_distinct_count",
			"SUM(timelaps)/SUM(total) as request_time_average",
		).
		Group("endpoint_id").
		Find(&data).
		Error(); err != nil {
		return nil, err
	}

	return data, nil
}
