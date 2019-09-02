package main

import (
  "fmt"
  "decred-roi-api/datab"
  "decred-roi-api/fetchapi"
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
  //testConnect, err := datab.DBTest()
  result1 := dbObj.ReadDB(1)
  result2 := dbObj.ReadDB(10)

  fmt.Println(result1)
  fmt.Println(result2)

  // Create new instance of the People struct
  fetchObj := fetchapi.People{}
  // Call the TestGet method for the People struct instance
  result3 := fetchObj.TestGet()

  fmt.Printf("%d people are currently in space\n", result3.HowMany)
  fmt.Printf("Example: %s is on the %s\n", result3.Persons[0].Name, result3.Persons[0].Craft)
}