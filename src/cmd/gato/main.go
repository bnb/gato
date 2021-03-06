package main

import (
	"log"
	"net/http"

	"sockethandler"

	"github.com/googollee/go-socket.io"
)

func main() {
	io, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	sockethandler.HandleServer(io)

	http.Handle("/socket.io/", io)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Escuchando en localhost:4286")
	log.Fatal(http.ListenAndServe(":4286", nil))
}
