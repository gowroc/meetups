package main

import (
	"flag"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"

	"github.com/gowroc/meetups/graphql/db"
)

func main() {
	query := flag.String("query", "", "GraphQL query")
	flag.Parse()

	mongo, err := db.NewMongo("localhost:37017", "graphql-example")
	if err != nil {
		log.Fatalf("failed to connect to Mongo: %v", err)
	}
	defer mongo.Close()

	psql, err := db.NewPostgres("localhost", 15432, "postgres", "graphql")
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}
	defer psql.Close()

	var userType = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type: graphql.String,
			},
			"admin": &graphql.Field{
				Type: graphql.Boolean,
			},
			"permissions": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					username := p.Source.(*db.User).Login

					permissions, err := mongo.GetUserPermissions(username)
					if err != nil {
						return nil, errors.Wrap(err, "failed to get permissions")
					}

					return permissions, nil
				},
			},
		},
	})

	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "rootQuery",
			Fields: graphql.Fields{
				"hello": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "world", nil
					},
				},
				"user": &graphql.Field{
					Type: userType,
					Args: graphql.FieldConfigArgument{
						"login": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						login := p.Args["login"].(string)

						u, err := psql.GetUser(login)
						if err != nil {
							return nil, errors.Wrap(err, "failed to get user")
						}

						return u, nil
					},
				},
			},
		},
	)

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "rootMutation",
		Fields: graphql.Fields{
			"addPermission": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"login": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"permission": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					login := p.Args["login"].(string)
					permission := p.Args["permission"].(string)

					if err := mongo.AddUserPermission(login, permission); err != nil {
						return nil, errors.Wrap(err, "failed to add permission")
					}

					u, err := psql.GetUser(login)
					if err != nil {
						return nil, errors.Wrap(err, "failed to fetch user")
					}

					return u, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery,
			Mutation: rootMutation,
		},
	)

	params := graphql.Params{
		Schema:        schema,
		RequestString: *query,
	}

	res := graphql.Do(params)
	if res.HasErrors() {
		log.Fatalf("Failed to get data from GraphQL: %v", res.Errors)
	}

	log.Println(res.Data)
}
