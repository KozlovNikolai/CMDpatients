package store

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(connStr string) {
	var err error
	//connStr := "postgres://dbuser:dbpass@localhost:35432/restapi_test?sslmode=disable"
	DB, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

}

func CloseDB() {
	DB.Close()
}
