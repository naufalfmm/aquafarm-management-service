package pondsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

func (u usecases) DeleteByID(ctx context.Context, id uint64, loginDeleted token.Data) error {
	if _, err := u.persistents.Repositories.Ponds.GetByID(ctx, id); err != nil {
		return err
	}

	return u.persistents.Repositories.Ponds.DeleteByID(ctx, id, loginDeleted)
}
