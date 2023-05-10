package model

import "github.com/graphql-go/graphql"

var (
	ProductInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "product",
		Description: "product is a information of the item",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
)
