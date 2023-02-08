package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (r repositories) GetAll(ctx context.Context, req dto.PondListRequest) (dao.Ponds, error) {
	var ponds dao.Ponds

	if err := req.Apply(r.resources.MySql.GetDB().WithContext(ctx)).
		Preload("Farm").
		Find(&ponds).
		Error(); err != nil {
		r.resources.Logger.Error(ctx, "error when getting all ponds",
			zapLog.SetAttribute("req", req),
			zapLog.SetAttribute("error", err),
		)
		return nil, err
	}

	return ponds, nil
}
