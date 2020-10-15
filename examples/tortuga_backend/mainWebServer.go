package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type webServer struct {
	app           *App
	instance      *http.Server
	activeSockets []*websocket.Conn
	addr          *string
	upgrader      websocket.Upgrader
}

func (ws *webServer) init(app *App) {
	ws.app = app
	ws.addr = flag.String("addr", ":8080", "http service address")
	ws.instance = &http.Server{Addr: *ws.addr}
	ws.upgrader = websocket.Upgrader{}
}

func (ws *webServer) run() {
	http.HandleFunc("/message", ws.socket)
	http.Handle("/", http.FileServer(http.Dir("../tortuga_frontend+/build/web")))

	log.Println(ws.instance.ListenAndServe())
	ws.app.waitInstance.Done()
}

func (ws *webServer) socket(w http.ResponseWriter, r *http.Request) {
	wsocket, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer wsocket.Close()

	// Don't allow websocket more than one
	if len(ws.activeSockets) > 0 {
		ws.activeSockets[0].Close()
		ws.activeSockets = ws.activeSockets[1:]
	}

	ws.activeSockets = append(ws.activeSockets, wsocket)

run:
	for {
		message := Message{}
		err := wsocket.ReadJSON(&message)
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

	log.Println(wsocket.RemoteAddr())
}
