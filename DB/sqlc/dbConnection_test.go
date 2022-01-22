package DB

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver            = "postgres"
	dbConnectionDetails = "postgresql://root:krish@knight8@localhost:5432/simple_transfers?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, connerr := sql.Open(dbDriver, dbConnectionDetails)
	if connerr != nil {
		log.Fatalln("Connection Failed, Error--> ", connerr)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}
