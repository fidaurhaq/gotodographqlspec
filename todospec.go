package todospec

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/segmentio/ksuid"
)

// Schema for graphql which describes the stucture of its queries, data and mutations
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {
		todo(id: ID!): Todo
		alltodos: [Todo]!
	}

	type Mutation {
		createTodo(label: String!, doneStatus: Boolean!): Todo
		updateTodo(id: ID!, label: String!, doneStatus: Boolean!) : Todo
		deleteTodo(id: ID!) : Todo
	}

	type Todo{
		id: ID!
		label: String!
		doneStatus: Boolean!
	}
`

type todo struct {
	ID         string
	Label      string
	DoneStatus bool
}

var todos = []*todo{
	{
		ID:         "1000",
		Label:      "Revise knowledge on pointers",
		DoneStatus: false,
	},
	{
		ID:         "1001",
		Label:      "Make a gotdamn amazing go graphql server for a generic todo app",
		DoneStatus: false,
	},
}

var todoData = make(map[string]*todo)

func init() {
	for _, t := range todos {
		todoData[t.ID] = t
	}
}

// Resolver used to
type Resolver struct{}

// Todo returns a single todo based on ID
func (r *Resolver) Todo(args struct{ ID string }) *todoResolver {
	if t := todoData[args.ID]; t != nil {
		return &todoResolver{t}
	}
	return nil
}

// Alltodos returns a list of todo items
func (r *Resolver) Alltodos() []*todoResolver {
	var tl []*todoResolver
	for _, atodo := range todoData {
		tl = append(tl, &todoResolver{atodo})
	}
	return tl
}

// CreateTodo appends list of current todos with a new todo with a new guid
func (r *Resolver) CreateTodo(args *struct {
	Label      string
	DoneStatus bool
}) *todoResolver {
	mytodo := &todo{
		Label:      args.Label,
		DoneStatus: args.DoneStatus,
	}
	mytodo.ID = ksuid.New().String()
	todos = append(todos, mytodo)
	for _, t := range todos {
		todoData[t.ID] = t
	}
	return &todoResolver{mytodo}
}

// UpdateTodo finds the todo using ID and updates it accordindly
func (r *Resolver) UpdateTodo(args *struct {
	ID         string
	Label      string
	DoneStatus bool
}) *todoResolver {
	mytodo := &todo{
		ID:         args.ID,
		Label:      args.Label,
		DoneStatus: args.DoneStatus,
	}

	for _, t := range todoData {
		if t.ID == args.ID {
			todoData[t.ID] = mytodo
		}
	}
	return &todoResolver{mytodo}
}

// DeleteTodo finds the todo using ID and deletes it
func (r *Resolver) DeleteTodo(args *struct {
	ID         string
	Label      string
	DoneStatus bool
}) *todoResolver {
	mytodo := &todo{}
	for _, t := range todoData {
		if t.ID == args.ID {
			mytodo = todoData[args.ID]
			delete(todoData, args.ID)
		}
	}
	return &todoResolver{mytodo}
}

type todoResolver struct {
	t *todo
}

func (r *todoResolver) ID() graphql.ID {
	return graphql.ID(r.t.ID)
}

func (r *todoResolver) Label() string {
	return r.t.Label
}

func (r *todoResolver) DoneStatus() bool {
	return r.t.DoneStatus
}
