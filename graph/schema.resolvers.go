package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"fmt"
	"go-graphql-mongodb-project/graph/model"
	"go-graphql-mongodb-project/database"

)
var db = database.Connect()
// CreateBoard is the resolver for the createBoard field.
func (r *mutationResolver) CreateBoard(ctx context.Context, input model.CreateBoardInput) (*model.Board, error) {
	// Call the database function to create the board
	return db.CreateBoard(ctx, input)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Boards is the resolver for the boards field.
func (r *queryResolver) Boards(ctx context.Context) ([]*model.Board, error) {
	return db.GetBoards(), nil
}

// In your resolver.go or equivalent file
func (r *queryResolver) BoardByID(ctx context.Context, id string) (*model.Board, error) {
	// Call your DB function to retrieve the board by ID
	return db.GetBoardByID(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	var db = database.Connect()
*/
