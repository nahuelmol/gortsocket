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

func TheWSconn(w http.ResponseWriter, r *http.Request){
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading")
    }
    defer conn.Close()


    for {
        msgType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("error reading message")
            break
        }

        log.Println("received message:", string(p))

        switch msgType {
        case websocket.TextMessage:
            log.Println("text message")
        case websocket.BinaryMessage:
            log.Println("binary message")
        }

        err = conn.WriteMessage(websocket.TextMessage, []byte("received by server"))
        if err != nil {
            log.Println("error sending message")
            break
        }
    }

}


