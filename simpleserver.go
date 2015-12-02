package main

import "fmt"
import "net/http"

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello, Web!")
    })

  http.ListenAndServe(":8080", nil)
}
