package chat

//Native Golang packages
import (
	"html/template"
	"net/http"
)

//Global chat html template file. Ensure it parses by calling template.Must.  Calls Panic if template does not render correctly.  Redirect to err page.
var homeTemplate = template.Must(template.ParseFiles("./chat/templates/home.html"))

//A room struct that will serve as a simple chat room.  A room has fields: connections, a register connection channel, an unregister connection channel, a message channel, and an outgoing channel.
type Room struct {
	// A room will have client-connections.  Connections is a map-field used to monitor connections
	connections map[*connection]bool

	// Channel of type byte slice. Slices are similar to dynamic arrays. Channels are like queues. They're used to communicate between goroutines(threads).
	outgoing chan []byte

	// This channel will be used to register connections to the connections map.  When there is a new connection it will be placed in the register channel to be placed in the map.
	register chan *connection

	// If a connection is inactive, it will be placed in the unregister channel, and dropped from the connection map.
	unregister chan *connection

  //Slice of messages currently in the chat room.  All messages are appended to the chat when a new connection is made.
	messages []byte
}

// Render the html template
func (*Room) RenderRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

// Makes and returns an initialized room struct.
func (*Room) AddRoom() Room {
	room := Room{
		broadcast:   make(chan []byte),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
		messages:    make([]byte, 0),
	}
	return room
}

// Method to run the room.  Makes it go into an infinite for loop, when a new connection connects to the chat, registers the connection, appends messages in message slice to the connection's chat
// if the connection sends messages, places the messages in the outgoing channel and sends to all connections in connection map.  When a connection isn't active, places it in the unregister channel
// to be deleted from the connection map.
func (R *Room) Run() {
	// This is a for select pattern.  The for is an infinite loop.  The select acts almost like a switch statement, and keeps checking each "case".
	// A case evaluates to true in a for select pattern when a channel has an item inside of it.
	for {
		select {
		// This is triggered when there is a connection that is waiting to be registered, in the register channel.
		case c := <-R.register:
		  // Add the connection c to the connections map
			R.connections[c] = true
		// When there is an inactive connection it is placed in the unregister channel.
		case c := <-R.unregister:
      // Delete the connection c from the map of connections
			if _, ok := R.connections[c]; ok {
				delete(R.connections, c)
				close(c.send)
			}
		// If there is message in channel outgoing, send to all connections in the connection map
		case m := <-R.outgoing:
			// Append an outgoing message to our messages slice
			m = append(m, []byte("\n")...)
			R.messages = append(R.messages, m...)
			// Loop over all connections currently in the connection map, grab each one, send the chat message to it.
			for c := range R.connections {
				select {
					case c.send <- m:
					default:
						close(c.send)
						delete(R.connections, c)
				}
			}
		}
	}
}

// The interesting part of this is how GO handles concurrency and parallellism.
//When running the program it would look something like this
func main() {
	// Create a new room
	room := Room.AddRoom();
	// Place it into the infinite for loop and make it run.
	// The go here in front of the function call, is placing the function in a goroutine.  In a typical program, execution would stall here waiting for the infinite for loop.
	// However the go keyword allows the function to run in parallel to the execution to the rest of the program.
	go room.Run();
	// In a similar fashion it is possible to create ten more rooms, all running in parallel
	for (i:=0; i<10; i++) {
		room := Room.AddRoom();
		go room.Run();
	}
  // In order to commmunicate with the room while in a go routine, a channel would be used, with a select.
}
