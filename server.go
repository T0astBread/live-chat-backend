package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"t0ast.cc/symflower-live-chat/db"
	"t0ast.cc/symflower-live-chat/graph"
	"t0ast.cc/symflower-live-chat/graph/generated"
	"t0ast.cc/symflower-live-chat/server"
)

const defaultPort = "8080"

//go:generate gqlgen generate
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.Reset()

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	srv := server.New(schema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
