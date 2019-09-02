package fetchapi

import (
  // Basic packages
  //"fmt"
  "io/ioutil"
  "log"
  // JSON fetch and handeling
  "net/http"
  "time"
  "encoding/json"
)

// People struct for the test api fetch
type People struct {
  Message string `json:"message"`
  Persons  []struct {
    Name  string `json:"name"`
    Craft string `json:"craft"`
  } `json:"people"`
  HowMany int `json:"number"`
}

// TestGet is a test fetch of public api json
func (peeps People) TestGet() People { // make new instance of People struct, called peeps
  /*========== Fetch and Unmarshal JSON data */
  url := "http://api.open-notify.org/astros.json"
  // declare the client that will be connecting, with timeout settings
  spaceClient := http.Client{
    Timeout: time.Second * 2, // Maximum of 2 secs
  }

  // delcare the http request as 'req'
  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.Fatal(err)
  }
  // Add a useragent header to the 'req'
  req.Header.Set("User-Agent", "fetch-test")

  // Perform the request, saving the response as res or the error as getErr
  res, getErr := spaceClient.Do(req)
  if getErr != nil {
    log.Fatal(getErr)
  }

  // Parse the Body of the response into the body var, or error to readErr
  body, readErr := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if readErr != nil {
    log.Fatal(readErr)
  }

  // Create a new instance of the people struct
  // ToDo: Move this out to main, or top, or its own init function
  //people1 := People{}
  // Unmarshal (ie parse) the json in body to the new people struct instance. Use & to reference the original, and not a copy
  jsonErr := json.Unmarshal(body, &peeps)
  if jsonErr != nil {
    log.Fatal(jsonErr)
  }

  return peeps
}