package app

import (
	"github.com/tanalam2411/go_websocket/pkg/common"
	"encoding/json"
	"log"
)

const sourceA1P1 = "App1Process1"


func App1Process1(channel common.Channel) {

	defer log.Printf("%s:: Dying", sourceA1P1)

	log.Printf("App1Process1:: Started reading from channel: %v", channel.Name)

	for {
		select {
		case msg := <- channel.TargetChannel.Channel:
			marshaledMsg, _ := json.Marshal(msg)
			log.Printf("%s:: Received message: %v", sourceA1P1, string(marshaledMsg))
		}
	}
}
