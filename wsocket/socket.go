package socket 

import ( 
    "log"
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader {
        ReadBufferSize: 1024,
        WriteBufferSize:1024,
}

func WsocketHandler(w http.ResponseWriter, r *http.Request){
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading")
    }
    defer conn.Close()
}


