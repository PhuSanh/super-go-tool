package cmd

import (
	"github.com/urfave/cli"
	"schedule-management/cron"
)

func cronAction(appContext *cli.Context) {
	cron.Run()
}

var Cron = cli.Command{
	Name: "cron",
	Usage: "Start cronjob",
	Action: func(appContext *cli.Context) {
		cronAction(appContext)
	},
}