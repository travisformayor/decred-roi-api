package main

import (
  "fmt"
  "decred-roi-api/datab"
  //"decred-roi-api/fetchapi"
)

func main() {
  dbObj := datab.Database{}
  dbObj.Connect()
  //testConnect, err := datab.DBTest()
  // result1 := dbObj.ReadDB(1)
  // result2 := dbObj.ReadDB(10)

  // fmt.Println(result1)
  // fmt.Println(result2)

  fmt.Println(dbObj.CheckUserTable())
  //dbObj.InsertRecord()
  dbObj.UpdateRecord(4)
  dbObj.DeleteRecord(2)

  // // Create new instance of the People struct
  // fetchObj := fetchapi.People{}
  // // Call the TestGet method for the People struct instance
  // result3 := fetchObj.TestGet()

  // fmt.Printf("%d people are currently in space\n", result3.HowMany)
  // fmt.Printf("Example: %s is on the %s\n", result3.Persons[0].Name, result3.Persons[0].Craft)
}