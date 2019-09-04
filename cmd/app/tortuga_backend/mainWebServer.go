package main

import (
	"log"
	"math"
	"net/http"

	"golang.org/x/net/websocket"
)

// Message ...
type Message struct {
	OriginalX int16 `json:"OriginalX"`
	OriginalY int16 `json:"OriginalY"`
	DraggedX  int16 `json:"DraggedX"`
	DraggedY  int16 `json:"DraggedY"`
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

	defer wsocket.Close()

	// Don't allow websocket more than one
	if len(ws.activeSockets) > 0 {
		ws.activeSockets[0].Close()
		ws.activeSockets = ws.activeSockets[1:]
	}

	ws.activeSockets = append(ws.activeSockets, wsocket)

	message := Message{}
	basicControl := BasicControl{0, 0, 0, 0}

run:
	for {
		err := websocket.JSON.Receive(wsocket, &message)
		if err != nil {
			log.Println("JSON decode error")
			break run
		} else {
			log.Println(
				message.OriginalX, message.OriginalY,
				message.DraggedX, message.DraggedY)

			// Don't go crazy!
			if math.Abs(float64(basicControl.OriginalX)) >= 600 ||
				math.Abs(float64(basicControl.OriginalY)) >= 600 ||
				math.Abs(float64(basicControl.DraggedX)) >= 600 ||
				math.Abs(float64(basicControl.DraggedY)) >= 600 {
				break run
			}

			basicControl.OriginalX = message.OriginalX
			basicControl.OriginalY = message.OriginalY
			basicControl.DraggedX = message.DraggedX
			basicControl.DraggedY = message.DraggedY

			ws.app.tortugaInstance.chanRequest <- basicControl
		}
	}

	log.Println(wsocket.Request().RemoteAddr)
}
