package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	conn *pgxpool.Pool
	*Queries
}

func (d *Database) Close() {
	d.conn.Close()
}

func GetDB(dbUrl string) (*Database, error) {
	log.Printf("Connecting to postgres")

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to create pool: %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Println("Error pinging database", err)
		panic(err)
	}
	return &Database{Queries: New(conn), conn: conn}, nil
}
