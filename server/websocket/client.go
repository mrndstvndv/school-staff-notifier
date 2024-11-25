package websocket

import (
	"bongserver/utils"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// WARN: This is a security risk if you don't check the origin!
			return true
		},
	}
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// INFO: specify that it is a protobuf message with a magic number
			m := []byte{utils.PROTOBUF_MAGIG}
			m = append(m, message...)

			c.conn.WriteMessage(websocket.BinaryMessage, m)
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {

	utils.LogDebug("Serving websocket")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.register <- client

	go client.writePump()
}
