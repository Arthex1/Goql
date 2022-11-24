package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"goql/database"
	"goql/graph/generated"
	"goql/graph/model"
	"goql/utils"
	"strings"
)

// User is the resolver for the user field.
func (r *mutationResolver) User(ctx context.Context, input model.NewUser) (*model.User, error) {






	password, hasherr := utils.HashPassword(input.Password)
	if hasherr != nil {
		return new(model.User), hasherr
	}

	usr, err := database.DB.Create_user(input.Name, input.BioText, input.Email, input.Developer, input.Plususer, *input.Links.Youtube, *input.Links.Twitter, *input.Links.Linkedin, *input.Links.Portfolio, *input.Links.Github, database.Convert_projects_to_db(input.Projects), password)
	if err != nil {
		return new(model.User), fmt.Errorf(err.Error())
	}
	to_ret := model.User{ID: usr.ID, Name: usr.Name, Bio: &model.Bio{Text: usr.Bio.Text, ID: usr.Bio.ID}, Email: usr.Email, Password: usr.Password, Links: &model.Links{ID: usr.Links.ID, Youtube: &usr.Links.Youtube, Twitter: &usr.Links.Youtube, Linkedin: &usr.Links.LinkedIN, Github: &usr.Links.Github, Portfolio: &usr.Links.Portfolio}, Badges: &model.Badges{ID: usr.Badges.ID, Developer: usr.Badges.Developer, Plususer: usr.Badges.PlusUser}, Projects: database.Convert_projects_to_query(usr.Projects)}
	return &to_ret, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, email *string, getBy string) (*model.User, error) {
	var ind *string
	getBy = strings.ToLower(getBy)  
	if getBy == "email" {ind = email } else if getBy == "id" { ind = id } else {	return new(model.User), fmt.Errorf("Invalid Indetifier, only email and id are supported.")	}

	usr, err := database.DB.Get_user(getBy, ind)
	if err != nil {
		return new(model.User), fmt.Errorf(err.Error())
	}
	to_ret := model.User{ID: usr.ID, Password: usr.Password, Name: usr.Name, Bio: &model.Bio{Text: usr.Bio.Text, ID: usr.Bio.ID}, Email: usr.Email, Links: &model.Links{ID: usr.Links.ID, Youtube: &usr.Links.Youtube, Twitter: &usr.Links.Youtube, Linkedin: &usr.Links.LinkedIN, Github: &usr.Links.Github, Portfolio: &usr.Links.Portfolio}, Badges: &model.Badges{ID: usr.Badges.ID, Developer: usr.Badges.Developer, Plususer: usr.Badges.PlusUser}, Projects: database.Convert_projects_to_query(usr.Projects)}
	return &to_ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
