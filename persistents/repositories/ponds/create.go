package pondsRepositories

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

func (r repositories) Create(ctx context.Context, pond dao.Pond) (dao.Pond, error) {
	if err := r.resources.MySql.GetDB().
		WithContext(ctx).
		Create(&pond).
		Error(); err != nil {
		return dao.Pond{}, err
	}

	return pond, nil
}
