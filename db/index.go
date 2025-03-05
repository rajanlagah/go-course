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

	// Create a connection pool configuration
	poolConfig, err := pgxpool.ParseConfig(config.Config.DbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse pool config: %v\n", err)
		os.Exit(1)
	}

	// Set pool configuration
	poolConfig.MaxConns = 50                       // Maximum number of connections
	poolConfig.MinConns = 10                       // Minimum number of connections
	poolConfig.MaxConnLifetime = 1 * time.Hour     // Maximum lifetime of a connection
	poolConfig.MaxConnIdleTime = 30 * time.Minute  // Maximum idle time for a connection
	poolConfig.HealthCheckPeriod = 1 * time.Minute // How often to check connection health

	// Create the connection pool
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DB, err = pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	// Test the connection pool
	err = DB.Ping(ctx)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return
	}

	log.Printf("Successfully connected to database with connection pool")
}
