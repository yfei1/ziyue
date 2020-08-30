package main


import (
  "net/http"
  "log"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

type Location struct {
  Name string
  Longtitude float32
  Latitude float32
}

func main() {
  resp, err := http.Get("http://localhost:8080/hello")
  if err != nil {
    log.Fatalf("fatal error when getting http response, desc: %v", err)
    return
  }

  defer resp.Body.Close()
  result, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(result))
  
  var b Location
  err = json.Unmarshal(result, &b)
  if err != nil {
    log.Fatalf("error decoding, desc: %s", err.Error())
  }
  fmt.Print(b.Name)
}
