package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // provides the mysql driver for sqlx
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// ClientConfig holds configuration information necessary to build a connection string.
type ClientConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

func getConnectionString(c ClientConfig) string {
	if c.Port == 0 {
		c.Port = 3306
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
	)
}

func pingDB(ctx context.Context, db *sqlx.DB) error {
	var target = struct {
		V int `db:"val"`
	}{}

	return db.GetContext(ctx, &target, "SELECT 1 AS val")
}

// Connect connects to a DB
func Connect(c ClientConfig) (*sqlx.DB, error) {
	connectionString := getConnectionString(c)

	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx open")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	if err := pingDB(ctx, db); err != nil {
		return nil, errors.Wrap(err, "sqlx ping")
	}

	return db, nil
}

// Load returns a ClientConfig for a db
func Load() (ClientConfig, error) {
	cfg := ClientConfig{}
	if err := envconfig.Process("", cfg); err != nil {
		logrus.Error(errors.Wrap(err, "envconfig: process"))
		return cfg, errors.Wrap(err, "envconfig: process")
	}

	return cfg, nil
}
