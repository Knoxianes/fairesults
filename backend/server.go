package main

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/graph"
	verifyHandler "Knoxiaes/fairesults/handlers/verify"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	h.AddTransport(&transport.Websocket{})

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

	r := gin.Default()
	database.InitDB()
	defer database.CloseDB()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/verify/:token", verifyHandler.GET)
	r.Run()
}
