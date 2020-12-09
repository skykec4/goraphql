package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skykec4/goraphql/graph/generated"
	"github.com/skykec4/goraphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprintf("%d", len(r.todos)+1),
		Text: input.Text,
		Done: input.Done,
		Date: input.Date,
	}
	r.todos = append(r.todos, todo)

	fmt.Println("넣었냐?", r.todos)
	fmt.Println("넣었냐?", todo)

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("찾았냐", r.todos)
	return r.todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	var data *model.Todo

	fmt.Println("들어옴 id :: '", id, "'")
	for i := 0; i < len(r.todos); i++ {
		fmt.Println("같니?1", &id, " :: ", r.todos[i].ID)

		if id == r.todos[i].ID {
			data = &model.Todo{
				ID:   r.todos[i].ID,
				Text: r.todos[i].Text,
				Done: r.todos[i].Done,
				Date: r.todos[i].Date,
			}
		}
	}

	return data, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
