package websocket

import (
	"github.com/tanalam2411/go_websocket/pkg/common"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

func WsServer(serverName string, pool *Pool, w http.ResponseWriter, r *http.Request) *Client{

	wsConn, wsErr := upgrader.Upgrade(w, r, nil)
	if wsErr != nil {
		log.Fatalf("%s:: Failed to upgrade: %v", serverName, wsErr)
		fmt.Fprintf(w, "%+v\n", wsErr)
	}

	client := &Client{
		ID:   guuid.New(),
		Conn: wsConn,
		Pool: pool,
		ServerName: serverName,
	}

	return client
}


func Start(serverName string, targets []common.Target, pool *Pool, client *Client) {

	pool.Register <- client
	client.Read(serverName, targets)

}
