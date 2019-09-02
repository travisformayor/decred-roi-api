package main

import (
	"fmt"
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
	dbObj := datab.Database{}
	dbObj.Connect()
	result := dbObj.ReadDB(1)
	result2 := dbObj.ReadDB(10)

	fmt.Println(result)
	fmt.Println(result2)

	//testConnect, err := datab.DBTest()
}