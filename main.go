package main

import (
	"database/sql"
	"log"

	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/krish8learn/simpleFootballTransfersRecorder/api"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver            = "postgres"
// 	dbConnectionDetails = "postgresql://root:krish@knight8@localhost:5432/simple_transfers?sslmode=disable"
// 	port                = "0.0.0.0:8080"
// )

func main() {
	config, err := Util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DbConnect, connerr := sql.Open(config.DBDriver, config.DBSource)
	if connerr != nil {
		log.Fatalln("Connection Failed, Error--> ", connerr)
	}

	transaction := DB.NewTransaction(DbConnect)
	server := api.NewServer(transaction)

	serverErr := server.Start(config.Port)

	if serverErr != nil {
		log.Fatalln("Cannot start server -->", serverErr)
	}
}
