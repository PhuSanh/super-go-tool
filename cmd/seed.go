package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"reflect"
	"schedule-management/database"
	"schedule-management/seed"
	"schedule-management/setting"
	"strings"
)

func seedAction(appContext *cli.Context) {
	setting.InitMysql()
	input := appContext.String("list")
	if input == "" {
		typeOfSeeder := reflect.TypeOf(seed.Seeder{})
		lenOfMethods := typeOfSeeder.NumMethod()
		fmt.Println("[1] You should pass argument --list all or -l all")
		fmt.Println("[2] Or you can pass one or more (use \",\" to separate) functions listed below:")
		for i := 0; i < lenOfMethods; i++ {
			fmt.Println("- " + typeOfSeeder.Method(i).Name)
		}
		return
	}

	seeder := seed.Seeder{
		DB: database.MysqlConn.DB,
	}
	s := reflect.ValueOf(seeder)
	var arguments []reflect.Value
	if input == "all" {
		lenOfMethods := s.NumMethod()
		for i := 0; i < lenOfMethods; i++ {
			s.Method(i).Call(arguments)
		}
		return
	}
	list := strings.Split(input, ",")
	for _, n := range list {
		s.MethodByName(n).Call(arguments)
	}
}

var Seed = cli.Command{
	Name: "seed",
	Usage: "Seed db",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "list, l",
			Usage: "List the name of seed func should be run",
		},
	},
	Action: func(appContext *cli.Context) error {
		seedAction(appContext)
		return nil
	},
}