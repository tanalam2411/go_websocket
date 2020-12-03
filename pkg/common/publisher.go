package common

import (
	"encoding/json"
	"errors"
	"fmt"
	gwebsocket "github.com/gorilla/websocket"
	"net/url"
)


type Target interface {
	GetTargetName() string
	WriteMessage(interface{}) error
}


type Channel struct {
	Name string
	TargetChannel MessageChannel
}

func(ch *Channel)GetTargetName() string {
	return ch.Name
}

func(ch *Channel)WriteMessage(msg interface{}) error{

	switch m := msg.(type) {
	case Message:
		ch.TargetChannel.Channel <- &m
	case *Message:
		ch.TargetChannel.Channel <- m
	case []byte:
		message := &Message{}
		err := json.Unmarshal(m, &message)
		if err != nil {
			return err
		}
		ch.TargetChannel.Channel <- message
	default:
		return errors.New("failed to write, message type not supported")
	}

	return nil
}


type WS struct {
	Scheme string
	Host string
	Path string
	Name string
}

func(ws *WS)GetTargetName() string {
	return ws.Name
}

func(ws *WS)URL()*url.URL{
	return &url.URL{
		Scheme:     ws.Scheme,
		Host:       ws.Host,
		Path:       ws.Path,
	}
}


func(ws *WS)Dial() (*gwebsocket.Conn, error) {
	fmt.Println(ws.URL().String())
	conn, _, err := gwebsocket.DefaultDialer.Dial(ws.URL().String(), nil)
	return conn, err
}

func(ws *WS)WriteMessage(msg interface{}) error {
	conn,  err := ws.Dial()
	if err != nil {
		return err
	}
	defer conn.Close()

	switch m := msg.(type) {
	case []byte:
		err = conn.WriteMessage(1, m)
		if err != nil {
			return err
		}
	case *Message:
		msg, err := json.Marshal(m)
		err = conn.WriteMessage(1, msg)
		if err != nil {
			return err
		}
	default:
		return errors.New("failed to write, message type not supported")
	}

	return nil
}



