package chat

import (
	"html/template"
  "net/http"
	"strconv"

	//chat imports
	"flag"
)

//chat globals
var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("./templates/home.html"))

//Global chat struct.  Ok global only ever should be one.
var cht = Chat {
  hubs : make([]*hub, 0),
}

type Chat struct {
  hubs []*hub
}

func (*Chat) addHub() {
  h := hub{
  	broadcast:   make(chan []byte),
  	register:    make(chan *connection),
  	unregister:  make(chan *connection),
  	connections: make(map[*connection]bool),
    messages:    make([]byte, 0),
  }
  cht.hubs = append(cht.hubs, &h)
}



//Runs all hubs in goroutines    Websocket needs to be embedded upon hub creation, not chat creation TODO
func (Chat) run() {
    for i, h := range cht.hubs{
      go h.run()
      num := strconv.Itoa(i)
      http.HandleFunc("/"+num, h.serveChat)
      http.HandleFunc("/ws"+"/"+num, h.serveWs)
    }
}

//Run function found in all applications to startup this module
func  (Chat) Run() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	flag.Parse()
	//Add 3 hubs for testing
	cht.addHub()
	cht.addHub()
	cht.addHub()
	//Start the controller struct
  cht.run()
}
