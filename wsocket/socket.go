package socket 

import ( 
    "log"
    "net/http"
    "encoding/json"
    
    "personal/wsservice/db"
    
    "github.com/gorilla/websocket"
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

        switch msgType {
        case websocket.TextMessage:
            json_string := string(p)
            var params map[string]interface{}
            json.Unmarshal([]byte(json_string), &params)

            username := params["username"].(string)
            id := params["id"].(string)
            access_key := params["access_key"].(string)

            log.Println("ckecking the %s\n data", username)
            db.AccessKeyInDB(access_key, id)

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


