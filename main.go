package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"schedule-management/cmd"
	"schedule-management/config"
	_ "schedule-management/config"
	"schedule-management/docs"
	"schedule-management/setting"
)

func init() {
	_ = setting.InitLogger()
	_ = setting.InitMysql()
	_ = setting.InitRedis()
}

func main() {

	setSwaggerConfig()

	app := cli.NewApp()

	app.Name = "schedule management"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		cmd.Seed,
		cmd.Migrate,
		cmd.Start,
		cmd.Test,
		cmd.Cron,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setSwaggerConfig() {
	var cfg = config.GetConfig()
	docs.SwaggerInfo.Title 			= fmt.Sprintf("%s API", cfg.Name)
	docs.SwaggerInfo.Description 	= " This is a document of Schedule Management server"
	docs.SwaggerInfo.Version 		= cfg.Version
	docs.SwaggerInfo.BasePath 		= fmt.Sprintf("/api/%s", cfg.ApiVersion)
	docs.SwaggerInfo.Host 			= fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}