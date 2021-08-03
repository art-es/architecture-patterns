package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	Database     string
	MaxIdleConns int
	MaxOpenConns int
}

func Connect(config Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Database)

	var counter int
	for {
		log.Println("[INFO] attempt to connect to database...")
		db, err := sql.Open("postgres", dataSourceName)
		if err == nil && db.Ping() == nil {
			db.SetMaxIdleConns(config.MaxIdleConns)
			db.SetMaxOpenConns(config.MaxOpenConns)
			return db, nil
		}

		if counter == 30 {
			log.Println("[ERROR] reached maximum number of attempt connecting to database")
			return nil, errors.New("database connection failed")
		}

		log.Printf("[WARN] attempt connecting to database failed, will be repeated in one second, details: %v\n", err)
		time.Sleep(time.Second)
		counter++
	}
}
