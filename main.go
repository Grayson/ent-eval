package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"entgo.io/ent/dialect"
	"github.com/Grayson/ent-eval/ent"
	"github.com/Grayson/ent-eval/ent/ogent"
	"github.com/Grayson/ent-eval/ent/todo"

	_ "github.com/mattn/go-sqlite3"
)

type handler struct {
	*ogent.OgentHandler
}

// Status implements ogent.Handler.
func (handler) ServerStatus(ctx context.Context) (ogent.ServerStatusRes, error) {
	status := ogent.ServerStatusOK{
		Data: strings.NewReader("We're okay"),
	}
	return &status, nil
}

func main() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	seed(client, ctx)

	server, err := ogent.NewServer(handler{
		OgentHandler: ogent.NewOgentHandler(client),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}

func seed(client *ent.Client, ctx context.Context) {
	parent := client.Todo.Create().
		SetText("Create seed data").
		SetStatus(todo.StatusCompleted).
		SetPriority(1).
		SaveX(ctx)

	client.Todo.CreateBulk(
		client.Todo.Create().
			SetText("child seed data 1").
			SetParent(parent),
		client.Todo.Create().
			SetText("child seed data 2").
			SetParent(parent),
		client.Todo.Create().
			SetText("child seed data 3").
			SetParent(parent),

		client.Todo.Create().
			SetText("seed data 1"),
	).SaveX(ctx)

}
