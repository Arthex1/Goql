package main

import (
	// "goql/graph/generated"
	// "log"
	// "net/http"
	"goql/helpers"
	"os"
	
	"goql/routes"
	"github.com/gin-gonic/gin"
	// "goql/resolvers"
	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default() 
	server.GET("/", helpers.PlaygroundHelper())
	server.POST("/query", helpers.GraphQLHelper())
	server.POST("/email", routes.Email)
	server.Static("/getimage", "./images")
	server.POST("/postimage/:userid", routes.PostImage)
	server.POST("/verifypassword", routes.VerifyPass) 
	
	server.Run()
	
	// 

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
