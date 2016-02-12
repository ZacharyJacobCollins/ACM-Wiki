package main

import (
  "net/http"
	"./wiki/Models"
	"./chat/Models"
)

func main() {
  http.ListenAndServe(":1337", nil)
}
