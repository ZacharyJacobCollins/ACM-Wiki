package main

import (
  "net/http"
  "/wiki"
  "/chat"
)

func main() {

  http.ListenAndServe(":1337", nil)
}
