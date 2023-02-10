package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/naufalfmm/aquafarm-management-service/migrations/model"
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func isTargetVersionExist(ctx context.Context, resources resources.Resources, targetVersion string) (bool, error) {
	if targetVersion == "" {
		return true, nil
	}

	var count int64
	if err := resources.MySql.WithContext(ctx).
		Model(&model.MigrationLog{}).
		Where("id", targetVersion).
		Count(&count).
		Error(); err != nil {
		return false, err
	}

	return count > 0, nil
}

func rollbackVersion(ctx context.Context, resources resources.Resources) (model.MigrationLog, error) {
	var (
		log model.MigrationLog
	)
	if err := resources.MySql.WithContext(ctx).
		Order("id DESC").
		Take(&log).
		Error(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.MigrationLog{}, nil
		}

		return model.MigrationLog{}, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return model.MigrationLog{}, err
	}

	sqlLocation := "/sql/"
	if !strings.Contains(sqlLocation, "migrations") {
		sqlLocation = "/migrations/sql/"
	}

	filePath := fmt.Sprintf("%s%s%s_%s_rollback.sql", wd, sqlLocation, log.ID, log.Name)
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
		if err := checkConnection(ctx.Context, resources); err != nil {
			return err
		}

		ver := ctx.String("version")

		isExist, err := isTargetVersionExist(ctx.Context, resources, ver)
		if err != nil {
			return err
		}

		if !isExist {
			return nil
		}

		resources.MySql.Begin()
		defer resources.MySql.Rollback()

		log, err := rollbackVersion(ctx.Context, resources)
		if err != nil {
			return err
		}
		for log.ID != ver && log.ID != "" {
			log, err = rollbackVersion(ctx.Context, resources)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func Rollback(resources resources.Resources) *cli.Command {
	return &cli.Command{
		Name:    "rollback",
		Usage:   "rollback --version <version>",
		Aliases: []string{"r"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "version",
				Aliases:  []string{"v"},
				Required: false,
			},
		},
		Action: rollback(resources),
	}
}
