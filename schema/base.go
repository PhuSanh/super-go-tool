package schema

import (
	"github.com/graphql-go/graphql"
	"schedule-management/resolver"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.GetUserByID,
			},
			"users": &graphql.Field{
				Type: graphql.NewList(User),
				Resolve: resolver.GetListUsers,
			},
		},
	})

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)