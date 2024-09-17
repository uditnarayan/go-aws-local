package main

import (
	"graphql/movies"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (query) Hello() string { return "Hello, world!" }

func main() {
	s := `
        type Query {
                hello: String!
        }
    `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})

	moviesHandler, err := movies.NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/movies", moviesHandler.RelayHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
