package main

import (
	"context"
	"os"

	"github.com/naufalfmm/aquafarm-management-service/migrations/cmd"
	"github.com/naufalfmm/aquafarm-management-service/migrations/resources"
	"github.com/urfave/cli/v2"
)

func main() {
	reso, err := resources.Init()
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.Migrate(reso),
		cmd.Rollback(reso),
	}

	if err = app.Run(os.Args); err != nil {
		reso.Logger.Error(context.Background(), "when running migration")
	}

	// fmt.Println(cmd.Migrate(context.Background(), reso))
	// fmt.Println(cmd.Rollback(context.Background(), reso, "1675911133918"))
}
