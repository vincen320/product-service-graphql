package model

import "github.com/graphql-go/graphql"

var (
	ClothType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "cloth",
		Description: "object of cloth",
		Fields: graphql.Fields{
			"mataerial": &graphql.Field{
				Type: graphql.String,
			},
			"size": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
)