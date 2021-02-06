package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"t0ast.cc/symflower-live-chat/db"
	"t0ast.cc/symflower-live-chat/graph/generated"
	"t0ast.cc/symflower-live-chat/graph/model"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.InsertUser(input)
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

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.GetUsers(), nil
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context) (<-chan *model.Message, error) {
	msgChan := make(chan *model.Message, 0)
	go func() {
		for {
			select {
			case <-ctx.Done():
				println("Closed messagePosted subscription")
				return
			default:
			}
			msgChan <- &model.Message{
				ID:      1,
				Content: "New message!",
				Poster:  dummyUser2,
			}
			time.Sleep(4 * time.Second)
		}
	}()
	return msgChan, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
