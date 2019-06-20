package graphql_handler

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"schedule-management/schema"
)

func Query(c echo.Context) (err error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: c.QueryParam("query"),
	})
	if len(result.Errors) > 0 {
		//return log.New("GraphQL ERROR")
	}
	return c.JSON(200, result)
}
