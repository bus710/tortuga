package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Message ...
type Message struct {
	Speed int16 `json:"Speed"`
	Angle int16 `json:"Angle"`
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
	// https://github.com/bus710/matrix2/blob/master/src/back/mainWebServer.go

	// Don't allow websocket more than one
	if len(ws.activeSockets) > 0 {
		ws.activeSockets[0].Close()
	}

	ws.activeSockets = append(ws.activeSockets, wsocket)

	message := Message{}
	speedAngle := SpeedAngle{0, 0}

run:
	for {
		err := websocket.JSON.Receive(wsocket, message)
		if err != nil {
			log.Println("JSON decode error")
		} else {
			log.Println(message.Speed, message.Angle)
			speedAngle.Speed = message.Speed
			speedAngle.Angle = message.Angle
		}

		if speedAngle.Speed >= 999 || speedAngle.Angle >= 999 {
			break run
		}

	}

	log.Println(wsocket.Request().RemoteAddr)
}
