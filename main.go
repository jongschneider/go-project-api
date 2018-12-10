package main

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/jongschneider/go-project/db"

	_ "github.com/go-sql-driver/mysql" // provides the mysql driver for sqlx
	"github.com/joho/godotenv"
	"github.com/jongschneider/go-project/server"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var cfg struct {
	Port int `envconfig:"PORT" required:"true" default:"3000"`

	MySQLDB   string `envconfig:"MYSQL_DB" required:"true" default:"project"`
	MySQLHost string `envconfig:"MYSQL_HOST" required:"true" default:"localhost"`
	MySQLPort int    `envconfig:"MYSQL_PORT" required:"true" default:"3306"`
	MySQLUser string `envconfig:"MYSQL_USER" required:"true" default:"root"`
	MySQLPass string `envconfig:"MYSQL_PW" default:""`

	Debug       bool   `envconfig:"DEBUG" default:"false"`
	Environment string `envconfig:"ENVIRONMENT" required:"true" default:"local"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Info(errors.Wrap(err, "godotenv"))
	}
	if err := envconfig.Process("", &cfg); err != nil {
		logrus.Error(errors.Wrap(err, "envconfig: process"))
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Info(errors.Wrap(err, "godotenv"))
	}
	// database, err := GetDB()
	server.New(cfg.Port).Start()
	log.Println("Application has run!")
}

// GetDB gets a sqlx DB
func GetDB() (*sqlx.DB, error) {
	dbConfig, err := db.Load()
	if err != nil {
		return nil, (errors.Wrap(err, "get dbConfig"))
	}
	database, err := db.Connect(dbConfig)
	if err != nil {
		return nil, (errors.Wrap(err, "connect to db"))
	}

	return database, nil
}
