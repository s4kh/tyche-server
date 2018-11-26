package types

import (
	"github.com/graphql-go/graphql"
)

var Envelope = graphql.NewObject(graphql.ObjectConfig{
	Name: "Envelope",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Number,
		},
	},
})
