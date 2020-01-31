package service

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/sinbadflyce/dictcrawler/generates"
	"github.com/sinbadflyce/dictcrawler/resolvers"
)

// Network ...
type Network struct {
}

const defaultPort = "3000"

// Listen ...
func (n *Network) Listen() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generates.NewExecutableSchema(generates.Config{Resolvers: &resolvers.LMResolver{}})))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
