package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"t0ast.cc/symflower-live-chat/graph"
	"t0ast.cc/symflower-live-chat/graph/generated"
)

const defaultPort = "8080"

type corsWrapper struct {
	underlying *handler.Server
}

// ServeHTTP wraps the underlying handler.Server and adds CORS
// headers
func (c *corsWrapper) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST")
		rw.Header().Set("Access-Control-Allow-Headers", "*")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.underlying.ServeHTTP(rw, req)
}

//go:generate gqlgen generate
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", &corsWrapper{srv})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
