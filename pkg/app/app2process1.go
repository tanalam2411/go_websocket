package app

import (
	"github.com/tanalam2411/go_websocket/pkg/common"
	"github.com/tanalam2411/go_websocket/pkg/websocket"
	"encoding/json"
	"fmt"
	"log"
)

const sourceA2P1 = "App2Process1"


func App2Process1(channel common.Channel, client *websocket.Client) {
	defer log.Printf("%s:: Dying", sourceA2P1)
	cfg := common.NewConfig("../config.yml")
	userClient := fmt.Sprintf("%s:%s", cfg.UserClient.Host, cfg.UserClient.Port)

	log.Printf("%s:: Started reading from channel: %v", sourceA2P1, channel.Name)

	for {
		select {
		case msg := <- channel.TargetChannel.Channel:
			marshaledMsg, _ := json.Marshal(msg)

			log.Printf("%s:: Received message: %v", sourceA2P1, string(marshaledMsg))

			respData := common.ResponseData{
				Topic:   "topic1",
				Message: map[string]int{"foo": 1, "bar": 2, "baz": 3},
			}

			marshaledMsg, _ = json.Marshal(respData)
			err := client.Conn.WriteMessage(1, marshaledMsg)
			if err != nil {
				log.Fatalf("%s::Failed to send msg to %s", sourceA2P1, err)
			}
			log.Printf("%s::Write msg to(User): %s/ws", sourceA2P1, userClient)
		}
	}

}
