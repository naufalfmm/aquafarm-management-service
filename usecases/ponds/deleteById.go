package pondsUsecases

import (
	"context"
)

func (u usecases) DeleteByID(ctx context.Context, id uint64, deletedBy string) error {
	if _, err := u.persistents.Repositories.Ponds.GetByID(ctx, id); err != nil {
		return err
	}

	return u.persistents.Repositories.Ponds.DeleteByID(ctx, id, deletedBy)
}
