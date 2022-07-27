package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sam-app/hackernews/graph/generated"
	"github.com/sam-app/hackernews/graph/model"
	"github.com/sam-app/hackernews/packages/tables"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link tables.Link
	link.Address = input.Address
	link.Title = input.Title

	linkID := link.Save()

	return &model.Link{ID: linkID, Title: link.Title, Address: link.Address}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	var post model.Post
	post.Title = input.Title
	post.Desc = input.Desc
	post.Content = input.Content

	postID, err := post.Save()
	if err != nil {
		return nil, err
	}

	return &model.Post{ID: postID, Title: post.Title, Desc: post.Desc, Content: post.Content}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	result, err := tables.GetAllLinks()
	for _, link := range result {
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address})
	}
	if err != nil {
		return nil, err
	}

	return resultLinks, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var resultPosts []*model.Post
	var post model.Post
	result, err := post.GetAllPosts()
	for _, post := range result {
		resultPosts = append(resultPosts, &model.Post{ID: post.ID, Title: post.Title, Desc: post.Desc, Content: post.Content})
	}
	if err != nil {
		return nil, err
	}

	return resultPosts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	result, err := post.GetPostById(id)
	fmt.Println(result)

	if err != nil {
		return nil, err
	}

	return &model.Post{ID: result.ID, Title: result.Title, Desc: result.Desc, Content: result.Content}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
