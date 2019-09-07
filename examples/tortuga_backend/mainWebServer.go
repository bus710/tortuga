package main

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

type webServer struct {
	app           *App
	instance      *http.Server
	activeSockets []*websocket.Conn
}

func (ws *webServer) init(app *App) {
	ws.app = app
	ws.instance = &http.Server{Addr: ":3000"}
}

func (ws *webServer) run() {
	http.Handle("/message", websocket.Handler(ws.socket))
	http.Handle("/", http.FileServer(http.Dir("../tortuga_frontend/build/web")))

	log.Println(ws.instance.ListenAndServe())
	ws.app.waitInstance.Done()
}

func (ws *webServer) socket(wsocket *websocket.Conn) {

	defer wsocket.Close()

	// Don't allow websocket more than one
	if len(ws.activeSockets) > 0 {
		ws.activeSockets[0].Close()
		ws.activeSockets = ws.activeSockets[1:]
	}

	ws.activeSockets = append(ws.activeSockets, wsocket)

	message := Message{}

run:
	for {
		err := websocket.JSON.Receive(wsocket, &message)
		if err != nil {
			log.Println("JSON decode error")
			break run
		} else {

			res := strings.Split(message.ButtonName, "/")
			if len(res) == 2 {
				if (res[0] == "forward" || res[0] == "none" || res[0] == "backward") &&
					(res[1] == "left" || res[1] == "none" || res[1] == "right") {

					ws.app.tortugaInstance.chanRequest <- BasicControl{res[0], res[1]}
				}
			}
		}
	}

	log.Println(wsocket.Request().RemoteAddr)
}
