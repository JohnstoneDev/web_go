/*
	how to work with websockets in Go using the gorilla/websocket package
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "websockets.html")

				fmt.Println("Server started")
		})

		http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)

			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to upgrade connection to WebSocket", http.StatusInternalServerError)
				return
			}

			for {
					// Read message from browser
					msgType, msg, err := conn.ReadMessage()
					if err != nil {
							return
					}

					// Print the message to the console
					fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

					// Write message back to browser
					if err = conn.WriteMessage(msgType, msg); err != nil {
							return
					}
			}
	})

	http.ListenAndServe(":8080", nil)
}