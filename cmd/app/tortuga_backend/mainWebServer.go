package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type webServer struct {
	app      *App
	instance *http.Server
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
}
