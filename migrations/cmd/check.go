package cmd

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/migrations/model"
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources"
)

func checkCreateMigrationLog(ctx context.Context, res resources.Resources) error {
	migrator := res.MySql.Migrator()

	if migrator.HasTable(model.MigrationLog{}) {
		return nil
	}

	return migrator.CreateTable(model.MigrationLog{})
}

func checkConnection(ctx context.Context, res resources.Resources) error {
	if err := res.MySql.Ping(); err != nil {
		return err
	}

	if err := checkCreateMigrationLog(ctx, res); err != nil {
		return err
	}

	return res.MySql.Ping()
}
