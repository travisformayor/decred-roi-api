package main

import (
	// "fmt"
	"decred-roi-api/datab"
)

// This method required adding a trust record to pg_hba.conf
// Is there a more secure way?
const (
  host     = "127.0.0.1"
  port     = 5432
  user     = "dcrtest"
  password = "testPassword"
  dbname   = "decred_test"
)

func main() {
	databaseObj := datab.Database{}
	databaseObj.DBTest()
	
	//testConnect, err := datab.DBTest()
}