// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/hiroto22/go-graphql-2/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Mutation struct {
		CreateTodo func(childComplexity int, input model.NewTodo) int
	}

	Query struct {
		Todos func(childComplexity int) int
	}

	Todo struct {
		Done func(childComplexity int) int
		ID   func(childComplexity int) int
		Text func(childComplexity int) int
		User func(childComplexity int) int
	}

	User struct {
		ID   func(childComplexity int) int
		Name func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.createTodo":
		if e.complexity.Mutation.CreateTodo == nil {
			break
		}

		args, err := ec.field_Mutation_createTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateTodo(childComplexity, args["input"].(model.NewTodo)), true

	case "Query.todos":
		if e.complexity.Query.Todos == nil {
			break
		}

		return e.complexity.Query.Todos(childComplexity), true

	case "Todo.done":
		if e.complexity.Todo.Done == nil {
			break
		}

		return e.complexity.Todo.Done(childComplexity), true

	case "Todo.id":
		if e.complexity.Todo.ID == nil {
			break
		}

		return e.complexity.Todo.ID(childComplexity), true

	case "Todo.text":
		if e.complexity.Todo.Text == nil {
			break
		}

		return e.complexity.Todo.Text(childComplexity), true

	case "Todo.user":
		if e.complexity.Todo.User == nil {
			break
		}

		return e.complexity.Todo.User(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputNewTodo,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

//go:embed "schema/schema.graphqls"
var sourcesFS embed.FS

func sourceData(filename string) string {
	data, err := sourcesFS.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("codegen problem: %s not available", filename))
	}
	return string(data)
}

var sources = []*ast.Source{
	{Name: "schema/schema.graphqls", Input: sourceData("schema/schema.graphqls"), BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
