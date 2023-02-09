package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/migrations/model"
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources"
	"github.com/urfave/cli/v2"
)

func generateName(splittedFileNames []string) string {
	var name string
	for _, splittedFileName := range splittedFileNames {
		name = fmt.Sprintf("%s%s_", name, splittedFileName)
	}

	return name[:len(name)-1]
}

func migrate(resources resources.Resources) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if err := checkConnection(ctx.Context, resources); err != nil {
			return err
		}

		resources.MySql.Begin()
		defer resources.MySql.Rollback()

		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		err = filepath.Walk(wd+"/sql/", func(p string, info fs.FileInfo, err error) error {
			_, file := filepath.Split(p)
			if file == "" {
				return nil
			}

			splittedFile := strings.Split(strings.TrimSuffix(file, filepath.Ext(file)), "_")

			if splittedFile[len(splittedFile)-1] != "migrate" {
				return nil
			}

			var count int64
			if err := resources.MySql.WithContext(ctx.Context).
				Model(model.MigrationLog{}).
				Where("id", splittedFile[0]).
				Count(&count).
				Error(); err != nil {
				return err
			}

			if count > 0 {
				return nil
			}

			content, err := ioutil.ReadFile(p)
			if err != nil {
				return err
			}

			if err := resources.MySql.WithContext(ctx.Context).
				Exec(string(content)).
				Error(); err != nil {
				return err
			}

			if err := resources.MySql.WithContext(ctx.Context).
				Create(model.MigrationLog{
					ID:        splittedFile[0],
					Name:      generateName(splittedFile[1 : len(splittedFile)-1]),
					MigrateAt: time.Now(),
				}).
				Error(); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		resources.MySql.Commit()

		return nil
	}
}

func Migrate(resources resources.Resources) *cli.Command {
	return &cli.Command{
		Name:    "migrate",
		Usage:   "migrate",
		Aliases: []string{"m"},
		Action:  migrate(resources),
	}
}
