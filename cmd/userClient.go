package main

import (
	"github.com/tanalam2411/go_websocket/pkg/common"
	"bufio"
	"encoding/json"
	"fmt"
	gwebsocket "github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
)

const source = "User"

var upgrader = gwebsocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func sendMsg(username, token string, target common.WS) {
	var loginMsg = common.Message{
		Type: gwebsocket.TextMessage,
		Body: common.LoginData{
			Username: username,
			Token: token,
		},
	}

	marshaledMsg, err := json.Marshal(loginMsg)
	log.Printf("%s::Sent message: %v to: %v/a2p1", source, string(marshaledMsg), target.Host)
	err = target.WriteMessage(marshaledMsg)
	if err != nil {
		log.Fatalf("%s:: Failed to Send msg: %v", source, err)
	}

}


func main() {

	cfg := common.NewConfig("../config.yml")
	serverApp2 := fmt.Sprintf("%s:%s", cfg.App2.Host, cfg.App2.Port)
	userClient := fmt.Sprintf("%s:%s", cfg.UserClient.Host, cfg.UserClient.Port)

	target := common.WS{
		Scheme: "ws",
		Host:   serverApp2,
		Path:   "/a2p1",
	}


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Username and Token separated by comma(,):")
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatalf("Error: ", err)
	}

	loginDetail := strings.Split(string(line), ",")
	username, token := loginDetail[0], loginDetail[1]
	log.Printf("Username: %s, token: %s", username, token)
	sendMsg(username, token, target)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsConn, wsErr := upgrader.Upgrade(w, r, nil)
		if wsErr != nil {
			log.Fatalf("%s:: Failed to upgrade: %v", source, wsErr)
			fmt.Fprintf(w, "%+v\n", wsErr)
		}

		defer func() {
			wsConn.Close()
		}()

		for {
			_, p, err := wsConn.ReadMessage()
			if err != nil {
				if gwebsocket.IsUnexpectedCloseError(err, gwebsocket.CloseGoingAway, gwebsocket.CloseAbnormalClosure) {
					log.Fatalf("%s:: Error occured while reading messages: %v", source, err)
				}
				break
			}

			log.Printf("%s:: Received msg: %s", source, string(p))

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter Username and Token seperated by comma(,):")
			line, _, err := reader.ReadLine()
			if err != nil {
				log.Fatalf("Error: ", err)
			}

			loginDetail := strings.Split(string(line), ",")
			username, token := loginDetail[0], loginDetail[1]
			log.Printf("Username: %s, token: %s", username, token)
			sendMsg(username, token, target)
		}

	})

	log.Fatal(http.ListenAndServe(userClient, nil))
}
