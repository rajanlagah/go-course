package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rajanlagah/go-course/config"
)

var DB *pgxpool.Pool

func InitDB() {
	var err error
	poolConfig, err := pgxpool.ParseConfig(config.Config.DbPath)
	if err != nil {
		log.Panic("Unable to parse config")
		os.Exit(1)
	}

	poolConfig.MaxConns = 45
	poolConfig.MinConns = 10
	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Hour
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	DB, err = pgxpool.NewWithConfig(context.Background(), poolConfig)

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
