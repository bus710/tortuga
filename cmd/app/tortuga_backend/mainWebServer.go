package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

// Message ...
type Message struct {
	ButtonName string `json:"ButtonName"`
}

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
	basicControl := BasicControl{2, 2}

run:
	for {
		err := websocket.JSON.Receive(wsocket, &message)
		if err != nil {
			log.Println("JSON decode error")
			break run
		} else {

			basicControl.x, err = strconv.Atoi(
				strings.Split(message.ButtonName, "/")[0])
			basicControl.y, err = strconv.Atoi(
				strings.Split(message.ButtonName, "/")[1])

			ws.app.tortugaInstance.chanRequest <- basicControl
		}
	}

	log.Println(wsocket.Request().RemoteAddr)
}
