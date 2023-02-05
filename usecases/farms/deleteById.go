package farmsUsecases

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

func (u usecases) DeleteByID(ctx context.Context, id uint64, loginDeleted token.Data) error {
	u.resources.MySql.StartTransaction(ctx)
	defer u.resources.MySql.RollbackTransaction(ctx)

	if _, err := u.persistents.Repositories.Farms.GetByID(ctx, id); err != nil {
		return err
	}

	if err := u.persistents.Repositories.Farms.DeleteByID(ctx, id, loginDeleted); err != nil {
		return err
	}

	if err := u.persistents.Repositories.Ponds.DeleteByFarmID(ctx, id, loginDeleted); err != nil {
		return err
	}

	u.resources.MySql.CommitTransaction(ctx)

	return nil
}
