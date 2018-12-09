package main

import (
	"log"

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
	// dbConfig, err := db.Load()
	// if err != nil {
	// 	log.Fatalln(errors.Wrap(err, "get dbConfig"))
	// }
	// s := server.New(port)
	// database, err := db.Connect(dbConfig)
	// if err != nil {
	// 	log.Fatalln(errors.Wrap(err, "connect to db"))
	// }
	s := server.New(cfg.Port)
	s.Start()
	log.Println("Application has run!")
}

// func getClient()
