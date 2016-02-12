package main

import (
  "net/http"
  "github.com/ZacharyJacobCollins/Wiki/chat"
)

func main() {
  c:=chat.NewChat()
  c.Run()
  http.ListenAndServe(":1337", nil)
}
