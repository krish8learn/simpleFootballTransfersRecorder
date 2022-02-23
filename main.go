package main

import (
	"database/sql"
	"log"

	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/api"

	_ "github.com/lib/pq"
)

const (
	dbDriver            = "postgres"
	dbConnectionDetails = "postgresql://root:krish@knight8@localhost:5432/simple_transfers?sslmode=disable"
	port                = "0.0.0.0:8080"
)

func main() {
	DbConnect, connerr := sql.Open(dbDriver, dbConnectionDetails)
	if connerr != nil {
		log.Fatalln("Connection Failed, Error--> ", connerr)
	}

	transaction := DB.NewTransaction(DbConnect)
	server := api.NewServer(transaction)

	serverErr := server.Start(port)

	if serverErr != nil {
		log.Fatalln("Cannot start server -->", serverErr)
	}
}
