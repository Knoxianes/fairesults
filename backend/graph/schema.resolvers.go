package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/handlers/graphqlHandlers"
	"context"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	return graphqlHandlers.CreateUser( input)
}

// CreateResult is the resolver for the createResult field.
func (r *mutationResolver) CreateResult(ctx context.Context, input model.NewResult) (*model.Result, error) {
	return graphqlHandlers.CreateResult(ctx, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return graphqlHandlers.Login(ctx, input)
}

// UpdatePassword is the resolver for the updatePassword field.
func (r *mutationResolver) UpdatePassword(ctx context.Context, input model.UpdatePassword) (bool, error) {
	return graphqlHandlers.UpdatePassword(ctx, input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser) (bool, error) {
	return graphqlHandlers.UpdateUser(ctx, input)
}

// UpdateResult is the resolver for the updateResult field.
func (r *mutationResolver) UpdateResult(ctx context.Context, input model.NewResult) (bool, error) {
	return graphqlHandlers.UpdateResult(ctx, input)
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, numberOfResults *int) (*model.User, error) {
	if numberOfResults == nil{
		return graphqlHandlers.User(1,0)
	}
	return graphqlHandlers.User(1,*numberOfResults)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
