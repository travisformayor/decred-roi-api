package main

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

const (
  host     = "127.0.0.1"
  port     = 5432
  user     = "dcrtest"
  password = "testPassword"
  dbname   = "decred_test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	fmt.Println("Hello. Test output: ", psqlInfo)

	// validate the arguments to open connection to db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ping the connection info to actually open the db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

  fmt.Println("Successfully connected!")
}