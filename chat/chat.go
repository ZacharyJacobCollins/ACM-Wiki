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

func NewChat() Chat{
  var chat = Chat {
    hubs : make([]*Hub, 0),
  }
  return chat
}

type Chat struct {
  hubs []*Hub
}

func (c *Chat) addHub() {
  h := Hub{
  	broadcast:   make(chan []byte),
  	register:    make(chan *connection),
  	unregister:  make(chan *connection),
  	connections: make(map[*connection]bool),
    messages:    make([]byte, 0),
  }
  c.hubs = append(c.hubs, &h)
}



//Runs all hubs in goroutines    Websocket needs to be embedded upon hub creation, not chat creation TODO
func (c *Chat) run() {
    for i, h := range c.hubs{
      go h.run()
      num := strconv.Itoa(i)
      http.HandleFunc("/"+num, h.serveChat)
      http.HandleFunc("/ws"+"/"+num, h.serveWs)
    }
}

//Run function found in all applications to startup this module
func  (c *Chat) Run() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	flag.Parse()
	//Add 3 hubs for testing
	c.addHub()
	c.addHub()
	c.addHub()
	//Start the controller struct
  c.run()
}
