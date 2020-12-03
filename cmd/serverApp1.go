package main

import (
	"github.com/tanalam2411/go_websocket/pkg/app"
	"github.com/tanalam2411/go_websocket/pkg/common"
	"github.com/tanalam2411/go_websocket/pkg/websocket"
	"fmt"
	"log"
	"net/http"
)


const a1p2ServerName = "App1Process2"

var(
	chan1App1 = common.MessageChannel{
		Channel: make(chan *common.Message),
	}
)


func a1p2WsRoutes(serverApp2 string) {
	targetC1A1 := common.Channel{Name: "Chan1/App1", TargetChannel:chan1App1}
	targetA2P3 := common.WS{
		Scheme: "ws",
		Host:   serverApp2,
		Path:   "/a2p3",
		Name: "App2Process3",
	}

	go app.App1Process1(targetC1A1)

	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/a1p2", func(w http.ResponseWriter, r *http.Request) {
		client := websocket.WsServer(a1p2ServerName, pool, w, r)
		websocket.Start(a1p2ServerName, []common.Target{&targetC1A1, &targetA2P3}, pool, client)
	})
}

func main() {
	cfg := common.NewConfig("../config.yml")
	serverApp1 := fmt.Sprintf("%s:%s", cfg.App1.Host, cfg.App1.Port)
	serverApp2 := fmt.Sprintf("%s:%s", cfg.App2.Host, cfg.App2.Port)
	a1p2WsRoutes(serverApp2)
	log.Printf("App1Process2:: Starting Server at address: '%s'", serverApp1)
	log.Fatal(http.ListenAndServe(serverApp1, nil))
}
