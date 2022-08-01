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
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user model.User
	user.Name = input.Name
	user.Username = input.Username
	user.Password = input.Password
	newUser, err := user.Save()
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	var post model.Post
	post.Title = input.Title
	post.Desc = input.Desc
	post.Content = input.Content
	post.UserID = "322bf063-cf70-48f8-9d95-d3fc29800bde" // TODO: get user id from context
	postID, err := post.Save()
	if err != nil {
		return nil, err
	}

	return &model.Post{ID: postID, Title: post.Title, Desc: post.Desc, Content: post.Content}, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	var err error
	post.ID = id
	post, err = post.Delete()
	if err != nil {
		return nil, err
	}
	fmt.Println("Deleted post", post)

	return &post, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return model.GetUserById(obj.UserID)
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
	fmt.Println("post", resultPosts)
	if err != nil {
		return nil, err
	}

	return resultPosts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	result, err := post.GetPostById(id)

	if err != nil {
		return nil, err
	}

	return &model.Post{ID: result.ID, Title: result.Title, Desc: result.Desc, Content: result.Content}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return model.GetUserById(id)
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]*model.Post, error) {
	return model.GetUserPosts(obj.ID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
