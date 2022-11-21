package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"goql/database"
	"goql/graph/generated"
	"goql/graph/model"
)

// User is the resolver for the user field.
func (r *mutationResolver) User(ctx context.Context, input model.NewUser) (*model.User, error) {
	usr, err := database.DB.Create_user(input.Name, input.BioText, input.Email)
	if err != nil {
		return new(model.User), fmt.Errorf(err.Error())
	}
	to_ret := model.User{ID: usr.ID, Name: usr.Name, Bio: &model.Bio{Text: usr.Bio.Text, ID: usr.Bio.ID, Email: usr.Bio.Email}}
	return &to_ret, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
