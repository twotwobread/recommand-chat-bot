package db

import (
	"context"

	_ "github.com/mattn/go-sqlite3"

	"recommand-chat-bot/internal/ent"
)

func InitInMemDB() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}
