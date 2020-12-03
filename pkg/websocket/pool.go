package websocket

import "log"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Printf("Pool:: Registered Client by ID: '%s', on server: '%s'", client.ID, client.ServerName)
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Printf("Pool:: UnRegistered Client by ID: '%s' from server: '%s'", client.ID, client.ServerName)
		}
	}
}
