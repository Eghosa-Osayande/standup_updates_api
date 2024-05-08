package database

import (
	"context"
	"fmt"
	"log"
	"os"

	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	// conn *pgxpool.Pool
	conn *pgx.Conn
	*Queries
}

func (d *Database) Close() {
	d.conn.Close(context.Background())
}

func GetDB() (*Database, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Africa/Lagos", host, user, pwd, name, port)
	log.Println("Connecting to postgres")

	ctx := context.Background()

	// conn, err := pgxpool.New(ctx, dbUrl)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Printf("Unable to connect to create db: %v", err)
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Println("Error pinging database", err)
		return nil, err
	}
	return &Database{Queries: New(conn), conn: conn}, nil
}
