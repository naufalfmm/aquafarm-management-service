package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (r repositories) GetAll(ctx context.Context, req dto.PondListRequest) (dao.Ponds, error) {
	var ponds dao.Ponds

	if err := req.Apply(r.resources.MySql.GetDB().WithContext(ctx)).
		Find(&ponds).Error(); err != nil {
		return nil, err
	}

	return ponds, nil
}
