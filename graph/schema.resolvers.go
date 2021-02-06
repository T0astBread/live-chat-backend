package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"t0ast.cc/symflower-live-chat/db"
	"t0ast.cc/symflower-live-chat/graph/generated"
	"t0ast.cc/symflower-live-chat/graph/model"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.InsertUser(input)
}

func (r *mutationResolver) PostMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	return db.InsertMessage(input)
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return db.GetMessages(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.GetUsers(), nil
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context) (<-chan *model.Message, error) {
	sub := make(chan *model.Message, 0)
	subHandle := db.RegisterMessageSubscription(sub)
	fmt.Printf("Registered message subscription %d\n", subHandle)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Un-registering message subscription %d...\n", subHandle)
			db.UnregisterMessageSubscription(subHandle)
		}
	}()
	return sub, nil
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
