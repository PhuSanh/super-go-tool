package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"net/http"
	"schedule-management/logger"
	"schedule-management/schema"
)

func Query(c echo.Context) (err error) {
	params := graphql.Params{Schema: schema.Schema, RequestString: c.QueryParam("query")}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		logger.Info(fmt.Sprintf("failed to execute graphql operation, errors: %+v", r.Errors), logger.LogFields{})
	}
	return c.JSON(http.StatusOK, r)
}