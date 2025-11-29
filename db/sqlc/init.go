package db

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"

	"github.com/yoru-no-mel/go-shippr/config"
)

func Connect(config config.Config) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}

func Close(conn *pgx.Conn) {
	err := conn.Close(context.Background())
	if err != nil {
		log.Fatalf("Unable to close connection: %v\n", err)
	}
}

func AutoMigrate(config config.Config) {

	path := fmt.Sprintf("file://%s", config.DBMigrationPath)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	m, err := migrate.New(path, dsn)
	if err != nil {
		log.Fatalf("unable to create migration: %v\n", err)
	}

	if config.DBRecreate {
		if err := m.Down(); err != nil {
			if err != migrate.ErrNoChange {
				log.Fatalf("unable to drop database: %v\n", err)
			}
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("unable to migrate database: %v\n", err)
	}
}

func Drop(config config.Config) {
	path := fmt.Sprintf("file://%s", config.DBMigrationPath)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	m, err := migrate.New(path, dsn)
	if err != nil {
		log.Fatalf("unable to create migration: %v\n", err)
	}
	if err := m.Down(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatalf("unable to drop database: %v\n", err)
		}
	}
}
