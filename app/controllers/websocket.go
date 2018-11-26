package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

// Websocket Controller
type Websocket struct {
	*revel.Controller
}

// Status returns provider status
func (c Websocket) Join(ws revel.ServerWebSocket) revel.Result {

	fmt.Println("join")

	if ws == nil {
		return nil
	}

	ws.MessageSend(receive)


	return nil
}

func receive() {
	fmt.Println("dsadasd")
	fmt.Println("dsadasd")
	fmt.Println("dsadasd")
	fmt.Println("dsadasd")
	fmt.Println("dsadasd")
}