package config

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	DATABASE = "jinpod_database"
)

type Config struct {
	MongoServer string
	Session     *mgo.Session
	Database    *mgo.Database
}

func InitConfig() *Config {
	var MONGOSERVER string
	if os.Getenv("MONGOSERVER") == "" {
		MONGOSERVER = "mongodb://opiumated:phoenix01@ds145283.mlab.com:45283/jinpod_database"
	}
	session, err := mgo.Dial(MONGOSERVER)

	if err != nil {
		log.Fatal("Database Error: ", err)
		os.Exit(2)
	}
	cfg := &Config{
		MongoServer: MONGOSERVER,
		Session:     session,
		Database:    session.DB(DATABASE),
	}
	return cfg
}

func NewConfig() *Config {
	return &Config{}
}
