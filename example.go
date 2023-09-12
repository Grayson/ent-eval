package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/Grayson/ent-eval/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Output:

	task1, err := client.Todo.
		Create().
		SetText("Add GraphQL Example").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}

	_, err = client.Todo.
		Create().
		SetText("Add Tracing Example").
		SetParent(task1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating task2: %v", err)
	}

	items, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		children, err := t.QueryChildren().All(ctx)
		fmt.Printf("%d: %q (%d children)\n", t.ID, t.Text, len(children))
		if err != nil {
			continue
		}
		for _, c := range children {
			fmt.Printf("-> %d\n", c.ID)
		}
	}
}
