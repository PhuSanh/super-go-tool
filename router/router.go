package router

import (
	"context"
	"fmt"
	"github.com/friendsofgo/graphiql"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"log"
	"os"
	"os/signal"
	"schedule-management/config"
	_ "schedule-management/docs"
	"schedule-management/gql"
	"schedule-management/handler"
	"syscall"
	"time"
)

var cfg = config.GetConfig()

func InitRouter() {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Any("/graphql", gql.Query)
	e.Any("/graphiql", echo.WrapHandler(graphiqlHandler))

	apiGroup := e.Group("/api")
	v1Group := apiGroup.Group("/v1")

	authHandler := new(handler.AuthHandler)
	v1Group.POST("/login", authHandler.Login)
	v1Group.POST("/check", authHandler.Check)

	userHandler := new(handler.UserHandler)
	v1Group.GET("/users", userHandler.GetList)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%v", cfg.Port)); err != nil {
			log.Println("⇛ shutting down the server")
			log.Printf("%v\n", err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
