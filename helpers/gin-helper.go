package helpers

import (
	"goql/graph/generated"
	"goql/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"
)

func PlaygroundHelper() gin.HandlerFunc {
	h := playground.Handler("Graphql Query", "/query")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request) 
	}
}

func GraphQLHelper() gin.HandlerFunc {
	srv :=  handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}})) 
	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}