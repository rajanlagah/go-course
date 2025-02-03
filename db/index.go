package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)
var DB *pgx.Conn

func InitDB() {
	urlExample := "postgres://postgres:adminPassword@localhost:5433/tasks"
	
	DB, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer DB.Close(context.Background()) // unmount
	log.Print(DB)
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Error in connecting DB")
		os.Exit(1)
	}

	log.Printf("Connected with DB")
}