package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rajanlagah/go-course/config"
)
var DB *pgx.Conn

func InitDB() {
	var err error
	DB, err = pgx.Connect(context.Background(), config.Config.DbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Error in connecting DB")
		os.Exit(1)
	}

	log.Printf("Connected with DB")
}