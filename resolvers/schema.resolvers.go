package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"goql/database"
	"goql/graph/generated"
	"goql/models"
	"math/rand"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	usrid := fmt.Sprintf("%v", rand.Int())
	skills := []*models.Skill{}
	for i := range input.Bio.Skils {
		skills = append(skills, (*models.Skill)(input.Bio.Skils[i]))
	}
	usr, errr := database.Database.CreateUser(&models.User{
		ID:   usrid,
		Name: input.Name,
		Bio: &models.Bio{
			URLCode: fmt.Sprintf("%v", rand.Int()),
			//UserID: usrid,
			Description: input.Bio.Description,
			Links: &models.Links{
				Github:    input.Bio.Links.Github,
				Twitter:   input.Bio.Links.Twitter,
				Youtube:   input.Bio.Links.Youtube,
				ID:        fmt.Sprintf("%v", rand.Int()),
				Portfolio: input.Bio.Links.Portfolio,
			},
			Skils: skills,
		},
	})

	if errr != nil {
		return usr, nil
	}
	return usr, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit *int) ([]*models.User, error) {
	usrs, err := database.Database.GetUsers(*limit) 
	if err != nil {
		return usrs, nil
	}
	return usrs, nil
}

// Bio is the resolver for the bio field.
func (r *queryResolver) Bio(ctx context.Context, urlCode string) (*models.Bio, error) {
	panic(fmt.Errorf("not implemented: Bio - bio"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
