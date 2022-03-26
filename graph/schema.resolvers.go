package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql/db"
	"go-graphql/graph/generated"
	"go-graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	entity, err := r.Db.Todo.CreateOne(
		db.Todo.Text.Set(input.Text),
		db.Todo.Done.Set(false),
		db.Todo.User.Link(db.User.ID.Equals(input.UserID)),
	).With(db.Todo.User.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	todo := &model.Todo{
		ID:   entity.ID,
		Text: entity.Text,
		Done: entity.Done,
		User: &model.User{
			ID:   entity.User().ID,
			Name: entity.User().Name,
		},
	}

	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	entity, err := r.Db.User.CreateOne(
		db.User.Name.Set(input.Name),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:   entity.ID,
		Name: entity.Name,
	}

	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	entities, err := r.Db.Todo.FindMany().With(db.Todo.User.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	todos := []*model.Todo{}
	for _, entity := range entities {
		todo := &model.Todo{
			ID:   entity.ID,
			Text: entity.Text,
			Done: entity.Done,
			User: &model.User{
				ID:   entity.User().ID,
				Name: entity.User().Name,
			},
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
