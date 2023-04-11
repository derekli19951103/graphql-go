package main

import (
	"fmt"
	"log"
	"os"

	"gql-go/db/model"
	resolver "gql-go/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
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
	h := playground.Handler("GraphQL", "/graphql")

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

	

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000","https://earth-rho.vercel.app"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	router.Use(cors.New(config),resolver.GinContextToContextMiddleware())
	
	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        panic(err)
    }
	err = db.AutoMigrate(&model.User{}, &model.Sketch{},&model.Session{}, &model.GeoObject{}, &model.Upload{})
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/graphql", graphqlHandler(db))
	router.GET("/", playgroundHandler())


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	router.Run(":"+port)
}
