package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
	src    = "postgres://myuser:mypass@localhost:5432/my_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	var err error
	testDB, err = sql.Open(driver, src)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
