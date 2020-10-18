package pg_db

import (
	"database/sql"
	"fmt"
	"github.com/cheremisin-sergey/forum/config"
	"time"
)

type Client struct {
	db *sql.DB
}

func New(config *config.Config) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "7444", config.DatabaseUsername, config.DatabasePassword, config.DatabaseName)

	fmt.Println(dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(20)

	return db, nil

	//return &Client{
	//	db: db,
	//}, nil
}

func (c *Client) Close() error {
	if c != nil && c.db != nil {
		return c.db.Close()
	}
	return nil
}
