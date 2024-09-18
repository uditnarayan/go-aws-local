package movies

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"graphql/movies/resolvers"
	"graphql/movies/schema"
)

type GraphqlHandler struct {
	Schema       *graphql.Schema
	RelayHandler *relay.Handler
	RootResolver *resolvers.RootResolver
}

func NewHandler() (*GraphqlHandler, error) {
	handler := &GraphqlHandler{}
	s, err := schema.String()
	if err != nil {
		return &GraphqlHandler{}, fmt.Errorf("error reading schema: %v", err)
	}
	db, err := resolvers.Connect()
	if err != nil {
		return &GraphqlHandler{}, fmt.Errorf("error connecting to database: %v", err)
	}
	handler.RootResolver = resolvers.NewRootResolver(db)
	options := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	handler.Schema = graphql.MustParseSchema(s, handler.RootResolver, options...)
	handler.RelayHandler = &relay.Handler{Schema: handler.Schema}
	return handler, nil
}
