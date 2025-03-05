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
	maxRetries := 5
	retryDelay := time.Second * 3

	config, err := pgxpool.ParseConfig(config.Config.DbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse pool config: %v\n", err)
		os.Exit(1)
	}

	// Adjust pool configuration for better performance under load
	config.MaxConns = 50 // Reduced to prevent overwhelming the DB
	config.MinConns = 10
	config.MaxConnLifetime = 15 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 30 * time.Second

	// Add connection timeout
	config.ConnConfig.ConnectTimeout = 5 * time.Second

	// Retry connection logic
	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		DB, err = pgxpool.NewWithConfig(ctx, config)
		cancel()

		if err != nil {
			log.Printf("Failed to create connection pool (attempt %d/%d): %v", i+1, maxRetries, err)
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			fmt.Fprintf(os.Stderr, "Unable to create connection pool after %d attempts: %v\n", maxRetries, err)
			os.Exit(1)
		}

		// Test the connection
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		err = DB.Ping(ctx)
		cancel()

		if err != nil {
			log.Printf("Failed to ping database (attempt %d/%d): %v", i+1, maxRetries, err)
			DB.Close()
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			fmt.Fprintf(os.Stderr, "Unable to connect to database after %d attempts: %v\n", maxRetries, err)
			os.Exit(1)
		}

		break
	}

	log.Printf("Successfully connected to DB with pool configuration: max=%d, min=%d connections", config.MaxConns, config.MinConns)
}
