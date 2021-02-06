package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

// CORSWrapper wraps an underlying handler.Server and adds CORS
// headers to every response
type CORSWrapper struct {
	underlying *handler.Server
}

// ServeHTTP wraps the underlying handler.Server and adds CORS
// headers
func (c *CORSWrapper) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "POST")
		res.Header().Set("Access-Control-Allow-Headers", "*")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.underlying.ServeHTTP(res, req)
}
