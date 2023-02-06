package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) GetByCode(ctx context.Context, code string) (dao.Pond, error) {
	var pond dao.Pond

	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Where("code", code).
		Take(&pond).
		Error(); err != nil {
		return dao.Pond{}, err
	}

	return pond, nil
}
