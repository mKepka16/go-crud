package initializers

import (
	"context"
	"log"
	"os"

	"github.com/mKepka16/go-crud/ent"
)

var DB *ent.Client

func ConnectToDB() {
	DB_URL := os.Getenv("DB_URL")

	var err error
	DB, err = ent.Open("postgres", DB_URL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
