package graph

import "github.com/felipefbs/graphql/internal/databases"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryRepository databases.CategoryRepository
}
