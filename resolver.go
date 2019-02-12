//go:generate go run scripts/gqlgen.go -v
package gqlgen_tutrial

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/tohutohu/gqlgen-tutrial/structs"
)

type Resolver struct {
	todos []gqlgen_tutrial.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (gqlgen_tutrial.Todo, error) {
	todo := gqlgen_tutrial.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]gqlgen_tutrial.Todo, error) {
	return r.todos, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *gqlgen_tutrial.Todo) (User, error) {
	return User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
