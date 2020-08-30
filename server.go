package main

import (
  "net/http"
  "fmt"
)

type Location struct {
  Name string
  Longtitude float32
  Latitude float32
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
    m := Location{
        Name: "Seattle",
        Longtitude: 123.1,
        Latitude: 321.1,
    }
    fmt.Fprint(w, m)
}

func main() {
  mux1 := http.ServerMux{}
  mux2 := http.ServerMux{}
  mux1.HandleFunc("/hello/", helloHandler)

  // Equivalent to http.DefaultServeMux.HandleFunc(...)
  http.HandleFunc("/hello/", helloHandler)
  http.ListenAndServe(":8080", nil)
}
