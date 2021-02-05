package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"t0ast.cc/symflower-live-chat/graph/generated"
	"t0ast.cc/symflower-live-chat/graph/model"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return dummyUser, nil
}

func (r *mutationResolver) PostMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	return &model.Message{
		ID:      1,
		Content: "This is my message",
		Poster:  dummyUser,
	}, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return []*model.Message{
		{
			ID:      1,
			Content: "This is my message",
			Poster:  dummyUser,
		},
		{
			ID:      2,
			Content: "Lorem ipsum",
			Poster:  dummyUser2,
		},
		{
			ID:      3,
			Content: "dolor sit amet",
			Poster:  dummyUser,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
