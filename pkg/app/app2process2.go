package app

import (
	"github.com/tanalam2411/go_websocket/pkg/common"
	"github.com/tanalam2411/go_websocket/pkg/websocket"
	"encoding/json"
	"fmt"
	"log"
)

const sourceA2P2 = "App2Process2"


func App2Process2(channel common.Channel) {
	defer log.Printf("%s:: Dying", sourceA2P2)

	log.Printf("%s:: Started reading from channel: %v", sourceA2P2, channel.Name)

	cfg := common.NewConfig("../config.yml")
	serverApp1 := fmt.Sprintf("%s:%s", cfg.App1.Host, cfg.App1.Port)

	for {
		select {
		case msg := <- channel.TargetChannel.Channel:
			marshaledMsg, err := json.Marshal(msg)
			log.Printf("%s:: Received message: %v", sourceA2P2, string(marshaledMsg))

			target := &common.WS{
				Scheme: "ws",
				Host:   serverApp1,
				Path:   "/a1p2",
			}
			err = websocket.Write(sourceA2P2, marshaledMsg, target)
			if err != nil {
				log.Fatalf("%s:: Failed to Send msg: %v", sourceA2P2, err)
			}
			log.Printf("%s:: Sending msg: %v, to server `App1Process2`: %v/a1p2", sourceA2P2, string(marshaledMsg), serverApp1)

		}
	}

}
