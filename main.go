package main

import (
	"yufengli/graphqlgo/api"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"errors"
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Define the GraphQL schema
    schema, _ := graphql.NewSchema(graphql.SchemaConfig{
        Query:    queryType, // Define the root Query type
        Mutation: mutationType, // Define the root Mutation type
    })

    // Create a new GraphQL handler with the schema
    graphqlHandler := handler.New(&handler.Config{
        Schema:   &schema,
        Pretty:   true,     
		GraphiQL: false,
		Playground: true,   
    })

    // Define a route for the GraphQL endpoint
    router.POST("/graphql", gin.WrapH(graphqlHandler))
	router.GET("/graphql", gin.WrapH(graphqlHandler))

    // Start the server
    router.Run(":8080")
}

var genshinWeaponType = graphql.NewObject(graphql.ObjectConfig{
    Name: "GenshinWeapon",
    Fields: graphql.Fields{
        "name": &graphql.Field{
            Type: graphql.String,
        },
        "rarity": &graphql.Field{
            Type: graphql.Int,
        },
        "level": &graphql.Field{
            Type: graphql.Int,
        },
    },
})


var genshinCharaterType = graphql.NewObject(graphql.ObjectConfig{
    Name: "GenshinCharater",
    Fields: graphql.Fields{
        "element": &graphql.Field{
            Type: graphql.String,
        },
        "actived_constellation_num": &graphql.Field{
            Type: graphql.Int,
        },
        "name": &graphql.Field{
            Type: graphql.String,
        },
        "level": &graphql.Field{
            Type: graphql.Int,
        },
        "rarity": &graphql.Field{
            Type: graphql.Int,
        },
        "weapon": &graphql.Field{
            Type: genshinWeaponType,
        },
    },
})

// Define the root Query type
var queryType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Query",
        Fields: graphql.Fields{
            "genshinCharaters": &graphql.Field{
                Type: graphql.NewList(genshinCharaterType),
				Args: graphql.FieldConfigArgument{
                    "uid": &graphql.ArgumentConfig{
                        Type: graphql.String,
                    },
                    "cookies":&graphql.ArgumentConfig{
                        Type: graphql.String,
                    },
                },
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    gc:=api.GenshinClient{Uid:p.Args["uid"].(string),Server_id:"os_usa", Cookies:p.Args["cookies"].(string)}
                    requestBody:= map[string]any{
                        "role_id":p.Args["uid"],
                        "server":"os_usa",
                    }
                    resp,err:=gc.Fetch("game_record/genshin/api/character",requestBody)
                    if err != nil {
                       return "", errors.New("can't get data") 
                    }     

                    return  resp.Data.Avatars, nil
                },
            },
        },
    },
)

// Define the root Mutation type
var mutationType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Mutation",
        Fields: graphql.Fields{
            "add": &graphql.Field{
                Type: graphql.Int,
                Args: graphql.FieldConfigArgument{
                    "x": &graphql.ArgumentConfig{
                        Type: graphql.Int,
                    },
                    "y": &graphql.ArgumentConfig{
                        Type: graphql.Int,
                    },
                },
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    x, _ := p.Args["x"].(int)
                    y, _ := p.Args["y"].(int)
                    return x + y, nil
                },
            },
        },
    },
)
