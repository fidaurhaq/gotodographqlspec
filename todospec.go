package todospec

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Schema for graphql which describes the stucture of its queries, data and mutations
var Schema = `
	schema {
		query: Query
	}

	type Query {
		todo(id: ID!): Todo
		alltodos: [Todo]!
	}

	type Todo{
		id: ID!
		label: String!
		doneStatus: Boolean!
	}
`

type todo struct {
	ID         graphql.ID
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

var todoData = make(map[graphql.ID]*todo)

func init() {
	for _, t := range todos {
		todoData[t.ID] = t
	}
}

// Resolver used to
type Resolver struct{}

// Todo returns a single todo based on ID
func (r *Resolver) Todo(args struct{ ID graphql.ID }) *todoResolver {
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

type todoResolver struct {
	t *todo
}

func (r *todoResolver) ID() graphql.ID {
	return r.t.ID
}

func (r *todoResolver) Label() string {
	return r.t.Label
}

func (r *todoResolver) DoneStatus() bool {
	return r.t.DoneStatus
}
