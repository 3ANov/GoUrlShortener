package db

import (
	"context"
	"fmt"
	"log"

	"GoUrlShortener/internal/utilities"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitDBPool() {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalf("unable to read configuration: %v", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Server, config.Port, config.Database)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("unable to parse database configuration: %v", err)
	}

	dbPool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func GetDBPool() *pgxpool.Pool {
	return dbPool
}

func CloseDBPool() {
	if dbPool != nil {
		dbPool.Close()
	}
}
