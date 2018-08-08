package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/googollee/go-socket.io"
)


func check(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {

	args := os.Args[1:]
	port := "5001"
	if len(args) > 0 {
		port = args[0]
	}

	server, err := socketio.NewServer(nil)
	check(err)
	server.On("connection", func(so socketio.Socket) {
		so.Join("main")
		so.On("tokencheck", func(token string) {

			events := make(chan *Event)
			go GetTokenData(token, events)

			for e := range events {
				so.Emit(e.Name, e.Data)
			}
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		fmt.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	fmt.Println("Serving at localhost:" + port + "...")
	http.ListenAndServe(":" + port, nil)
}