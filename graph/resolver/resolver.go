package resolver

import "github.com/skykec4/goraphql/graph/model"

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

//go:generate go run github.com/99designs/gqlgen

//Resolver ...
type Resolver struct {
	todos []*model.Todo
}
