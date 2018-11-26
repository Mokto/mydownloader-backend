

package websockets

import (
	"encoding/json"
	"localserver/app/models"
	"golang.org/x/net/websocket"
)

var connections = make(map[int]*websocket.Conn);

func AddConnection(ws *websocket.Conn, connId int) {
	connections[connId] = ws
}

func RemoveConnection(connId int) {
	delete(connections, connId)
}

func SendMessage(id string, message string) {
	for _, ws := range connections {
		websocketEvent := models.WebsocketEvent{ID: id, Data: message}
		data, err := json.Marshal(websocketEvent)
		if (err != nil) {
			return
		}
		websocket.Message.Send(ws, string(data))
	}
}