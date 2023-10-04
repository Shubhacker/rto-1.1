package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"
	"rto/graph/model"
	"rto/logic"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input *model.SignUpRequest) (*model.SignUpResponse, error) {
	// panic(fmt.Errorf("not implemented: SignUp - signUp"))
	response, err := logic.SignUpUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (*model.LoginResponse, error) {
	return logic.Login(ctx, input)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
