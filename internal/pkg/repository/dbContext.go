package repository

import (
	"context"
	"flag"
	"fmt"
	"go_template/config"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {
	configPath := flag.String("config", "config/config.yml", "path to config file")
	flag.Parse()
	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatal("cant get config", err)
	}
	databaseURL := cfg.DB.Ivalue

	dbpool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		if err != nil {
			return nil, err
		}
		os.Exit(1)
	}
	if dbpool != nil {
		log.Println("Database pool created successfully")
	} else {
		log.Println("Failed to create database pool")
	}
	return dbpool, nil
}
