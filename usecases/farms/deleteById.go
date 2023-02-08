package farmsUsecases

import (
	"context"
)

func (u usecases) DeleteByID(ctx context.Context, id uint64, deletedBy string) error {
	if _, err := u.persistents.Repositories.Farms.GetByID(ctx, id); err != nil {
		return err
	}

	u.resources.MySql.StartTransaction(ctx)
	defer u.resources.MySql.RollbackTransaction(ctx)

	if err := u.persistents.Repositories.Farms.DeleteByID(ctx, id, deletedBy); err != nil {
		return err
	}

	if err := u.persistents.Repositories.Ponds.DeleteByFarmID(ctx, id, deletedBy); err != nil {
		return err
	}

	u.resources.MySql.CommitTransaction(ctx)

	return nil
}
