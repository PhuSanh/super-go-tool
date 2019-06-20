package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"schedule-management/database"
)

var Test = cli.Command{
	Name: "test",
	Usage: "Test function",
	Action: func(appContext *cli.Context) {
		fmt.Println("test redis")
		err := database.RedisConn.Set("test-redis-golang", "okla", 100)
		fmt.Println(err)
		a, err := database.RedisConn.Get("test-redis-golang")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(a)
	},
}
