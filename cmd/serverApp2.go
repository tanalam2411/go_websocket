package main

import (
	"github.com/tanalam2411/go_websocket/pkg/app"
	"github.com/tanalam2411/go_websocket/pkg/common"
	"github.com/tanalam2411/go_websocket/pkg/websocket"
	"fmt"
	"log"
	"net/http"
)

const serverName = "ServerApp2"

var(
	chan1App2 = common.MessageChannel{
		Channel: make(chan *common.Message),
	}
	chan3App2 = common.MessageChannel{
		Channel: make(chan *common.Message),
	}
)


func app2WsRoutes() {
	targetC1A2 := common.Channel{Name: "Chan1/App2", TargetChannel:chan1App2}
	targetC3A2 := common.Channel{Name: "Chan3/App2", TargetChannel:chan3App2}
	go app.App2Process2(targetC1A2)


	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/a2p1", func(w http.ResponseWriter, r *http.Request) {
		client := websocket.WsServer("ServerApp2P1", pool, w, r)
		go app.App2Process1(targetC3A2, client)
		websocket.Start("ServerApp2P1", []common.Target{&targetC1A2}, pool, client)
	})

	http.HandleFunc("/a2p3", func(w http.ResponseWriter, r *http.Request) {
		client := websocket.WsServer("ServerApp2P3", pool, w, r)
		websocket.Start("ServerApp2P3", []common.Target{&targetC3A2}, pool, client)
	})
}

func main() {
	cfg := common.NewConfig("../config.yml")
	serverApp2 := fmt.Sprintf("%s:%s", cfg.App2.Host, cfg.App2.Port)
	app2WsRoutes()
	log.Printf("%s:: Starting Server at address: '%v'", serverName, serverApp2)
	log.Fatal(http.ListenAndServe(serverApp2, nil))
}
