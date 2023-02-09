package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/naufalfmm/aquafarm-management-service/migrations/model"
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources"
	"github.com/urfave/cli/v2"
)

func rollbackVersion(ctx context.Context, resources resources.Resources, wd string) (model.MigrationLog, error) {
	var (
		log model.MigrationLog
	)
	if err := resources.MySql.WithContext(ctx).
		Order("id DESC").
		Take(&log).
		Error(); err != nil {
		return model.MigrationLog{}, err
	}

	filePath := fmt.Sprintf("%s/sql/%s_%s_rollback.sql", wd, log.ID, log.Name)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return model.MigrationLog{}, err
	}

	if err := resources.MySql.WithContext(ctx).
		Exec(string(content)).
		Error(); err != nil {
		return model.MigrationLog{}, err
	}

	if err := resources.MySql.WithContext(ctx).
		Where("id", log.ID).
		Delete(model.MigrationLog{}).
		Error(); err != nil {
		return model.MigrationLog{}, err
	}

	return log, nil
}

func rollback(resources resources.Resources) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		ver := ctx.String("version")

		if err := checkConnection(ctx.Context, resources); err != nil {
			return err
		}

		resources.MySql.Begin()
		defer resources.MySql.Rollback()

		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		log, err := rollbackVersion(ctx.Context, resources, wd)
		if err != nil {
			return err
		}
		for log.ID != ver {
			log, err = rollbackVersion(ctx.Context, resources, wd)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func Rollback(resources.Resources) *cli.Command {
	return &cli.Command{
		Name:    "rollback",
		Usage:   "rollback --version <version>",
		Aliases: []string{"r"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "version",
				Aliases:  []string{"v"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error { return nil },
	}
}
