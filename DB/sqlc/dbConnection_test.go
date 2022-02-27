package DB

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver            = "postgres"
// 	dbConnectionDetails = "postgresql://root:krish@knight8@localhost:5432/simple_transfers?sslmode=disable"
// )

var testQueries *Queries
var testDbConnect *sql.DB

func TestMain(m *testing.M) {
	configs, configErr := Util.LoadConfig("../..")
	if configErr != nil {
		log.Fatal("cannot load config:", configErr)
	}
	var connerr error
	testDbConnect, connerr = sql.Open(configs.DBDriver, configs.DBSource)
	if connerr != nil {
		log.Fatalln("Connection Failed, Error--> ", connerr)
	}

	testQueries = New(testDbConnect)

	os.Exit(m.Run())

}
