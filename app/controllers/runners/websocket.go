package runners


import (
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	"localserver/app/services/websockets"
	"localserver/app/utils/links"
	"github.com/satori/go.uuid"
)

func StartWebsockets() {
	http.Handle("/", websocket.Handler(func (ws *websocket.Conn) {
		defer ws.Close()
		fmt.Println("Client Connected")

		connId := uuid.Must(uuid.NewV4()).String()
	
		websockets.AddConnection(ws, connId)
		go links.ListAndSend()
	
	
		for {
			var message string
			err := websocket.Message.Receive(ws, &message)
			if err != nil {
				websockets.RemoveConnection(connId)
				break
			}
		}
	}))

	fmt.Println("Websocket server is listening to : 9001")
	go http.ListenAndServe(":9001", nil)
}
