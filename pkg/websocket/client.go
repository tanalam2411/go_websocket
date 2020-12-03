package websocket

import (
	"github.com/tanalam2411/go_websocket/pkg/common"
	"encoding/json"
	"github.com/google/uuid"
	gwebsocket "github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   uuid.UUID
	Conn *gwebsocket.Conn
	Pool *Pool
	ServerName string
}


func (c *Client) Read(serverName string, targets []common.Target) {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			if gwebsocket.IsUnexpectedCloseError(err, gwebsocket.CloseGoingAway, gwebsocket.CloseAbnormalClosure) {
				log.Fatalf("%s:: Error occured while reading messages: %v", serverName, err)
			}
			break
		}

		log.Printf("%s:: Received msg: %s", serverName, p)

		message := &common.Message{}
		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Printf("%s:: Failed to unmarshal msg: %v", serverName, err)
		}

		for _, target := range targets{
			err := target.WriteMessage(message)
			if err != nil {
				log.Fatalf("%s:: Failed to send msg to target: %v, err: %v", serverName, target.GetTargetName(), err)
			}
			msg, err := json.Marshal(message)
			if err != nil {
				log.Fatalf("%s:: Failed to marshal msg: %v, for target: %v, err: %v", serverName, message, target.GetTargetName(), err)
			}
			log.Printf("%s:: Sending msg: %v, to target: %v", serverName, string(msg), target.GetTargetName())
		}
	}
}



func Write(source string, msg []byte, target common.Target) error {

	err := target.WriteMessage(msg)
	if err != nil {
		log.Fatalf("%s:: Write:: Failed to Send msg: %v", err)
		return err
	}
	return nil
}
