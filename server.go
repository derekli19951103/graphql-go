package main

import (
	"log"
	"os"

	"gql-go/db/model"
	resolver "gql-go/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"


// Defining the Graphql handler
func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{Resolvers: &resolver.Resolver{DB: db}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router.Use(resolver.GinContextToContextMiddleware())

	dsn := "host=192.168.2.117 user=postgres password=derekli dbname=test port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        panic(err)
    }
	err = db.AutoMigrate(&model.User{}, &model.Sketch{},&model.Session{}, &model.GeoObject{})
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/query", graphqlHandler(db))
	router.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	router.Run(":"+port)
}
