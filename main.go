package main

import (
  "net/http"
  "github.com/ZacharyJacobCollins/Wiki/chat"
  "github.com/ZacharyJacobCollins/Wiki/wiki"
)

func main() {

  http.ListenAndServe(":1337", nil)
}
