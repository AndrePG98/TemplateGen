package database

import (
	"TemplateGen/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Init(config *config.Config, maxAttempts int, retryInterval time.Duration) (*sql.DB, error) {
	var conn *sql.DB
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.USER, config.PASSWORD, config.HOST, config.DB_PORT, config.DB,
	)

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		conn, err = sql.Open("mysql", dsn)
		if err == nil {
			err = conn.Ping()
			if err == nil {
				conn.SetMaxIdleConns(10)
				conn.SetMaxOpenConns(100)
				conn.SetConnMaxIdleTime(time.Minute * 30)
				break
			}
		}
		time.Sleep(retryInterval)
	}

	return conn, err
}
