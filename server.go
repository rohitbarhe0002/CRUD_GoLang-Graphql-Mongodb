package main

import (
	"go-graphql-mongodb-project/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Learnings : 

// To generate schemas and genrated files:  go run github.com/99designs/gqlgen generate
// To run go server : go run server.go
// go generate ./...
// To install go modules : go mod tidy
// initialize new go project : go run github.com/99designs/gqlgen init

// create new dir and run this  above comand first 

// ref link for go : https://medium.com/the-godev-corner/build-efficient-apis-in-go-with-graphql-and-gqlgen-d8cbcb5edcfa
