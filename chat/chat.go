package chat

import (
	"net/http"
	"strconv"

	//chat imports
	"flag"
)

func NewChat() Chat {
	var chat = Chat{
		hubs: make([]*Hub, 0),
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

//Run function found in all applications to startup this module  TODO initialize with number of hubs to run.
//TODO on addHub needs to be the embed/run thing.  Needs to be by convo name as apposed to number, and TODO needs to be verification that it doesn't already exist

func (c *Chat) Run() {
	flag.Parse()
	//Add 3 hubs for testing
	c.addHub()
	c.addHub()
	c.addHub()
	//Start each hub in a goroutine
	for i, h := range c.hubs {
		go h.run()
		num := strconv.Itoa(i)
		http.HandleFunc("/"+num, h.executeHub)
		http.HandleFunc("/ws"+"/"+num, h.serveWs)
	}

}
