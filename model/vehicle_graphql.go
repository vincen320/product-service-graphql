package model

import "github.com/graphql-go/graphql"

var (
	VehicleType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "vehicle",
		Description: "object of vehicle",
		Fields: graphql.Fields{
			"engine": &graphql.Field{
				Type: graphql.String,
			},
			"wheel": &graphql.Field{
				Type: graphql.Int,
			},
		},
		Interfaces: []*graphql.Interface{
			ProductInterface,
		},
	})
)
