package socket

import (
	"bytes"
	"github.com/Peterliang233/debate/config"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var (
	writeWait = config.SocketSetting.WriteWait
	pongWait =  config.SocketSetting.PongWait
	pingPeriod = (pongWait * 9)/ 10
	maxMessageSize = config.SocketSetting.MaxMessageSize
)

var (
	newLine = []byte{'\n'}
	space = []byte{' '}
)

var upgrade = websocket.Upgrader{
	ReadBufferSize: config.SocketSetting.ReadBufferSize,
	WriteBufferSize: config.SocketSetting.WriteBufferSize,
}

type Client struct {
	hub *Hub
	coon *websocket.Conn
	send chan []byte

	username []byte
	roomID []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.coon.Close()
	}()

	c.coon.SetReadLimit(maxMessageSize)
	c.coon.SetReadDeadline(time.Now().Add(pongWait))
	c.coon.SetPongHandler(func(string) error {
		c.coon.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.coon.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
		c.hub.broadcast <- message
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.coon.Close()
	}()

	for {
		select {
		case message, ok := <- c.send:
			c.coon.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.coon.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.coon.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			n := len(message)
			for i := 0; i < n; i++ {
				w.Write(newLine)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil{
				log.Printf("error: %v",err)
				return
			}
		case <- ticker.C:
			c.coon.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.coon.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
		}
	}
}



func ServeWs(hub *Hub, c *gin.Context) {
	var req model.DebateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.ErrParameter,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrParameter],
			},
		})
		return
	}

	userName := req.UseName
	roomID := req.RoomID
	var upgrader = websocket.Upgrader{}

	coon, err := upgrader.Upgrade(c.Writer, c.Request, nil )
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub: hub,
		coon: coon,
		send: make(chan []byte, 256),
		username: []byte(userName),
		roomID: []byte(roomID),
	}

	client.hub.register <- client

	go client.WritePump()
	go client.ReadPump()

}