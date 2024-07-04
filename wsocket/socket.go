package socket 

import ( 
    "log"
    "net/http"
    "encoding/json"
    "sync/atomic"
    "fmt"
    "sync"
    
    "personal/wsservice/obj"
    "github.com/gorilla/websocket"
)

var Mutex sync.Mutex
var Clients = make(map[*websocket.Conn]bool)
var UserRegister = make(map[uint32]*obj.StackLocation)
var HmanyConnections int32

var UPGRADER = websocket.Upgrader {
        ReadBufferSize: 1024,
        WriteBufferSize:1024,
}

func TheWSconn(w http.ResponseWriter, r *http.Request){
    conn, err := UPGRADER.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading")
    }
    contextValue := r.Context().Value("v")
    fmt.Println("context value:", contextValue)

    defer conn.Close()
    Mutex.Lock()
    Clients[conn] = true
    Mutex.Unlock()
    for {
        msgType, p, err := conn.ReadMessage()
        if err != nil {
            Mutex.Lock()
            delete(Clients, conn)
            Mutex.Unlock()
            fmt.Println(err)
            break
        }
        atomic.AddInt32(&HmanyConnections, 1)

        switch msgType {
        case websocket.TextMessage:
            json_string := string(p)
            var params map[string]interface{}
            json.Unmarshal([]byte(json_string), &params)

            _, source_exists := params["source"]
            if !source_exists {
                rawMessage := RawMessage { content:string(p) }
                fmt.Println(rawMessage)
                message := "there's not a source"
                err = conn.WriteMessage(websocket.TextMessage, []byte(message))
            } else {
                result, err := StartSessionMsg(params)
                if err != nil {
                    fmt.Println("any session started")
                    break
                }
                fmt.Println(result)
                break
            }

            _, urequest := params["urequest"] //that's what a driver receives
            if urequest {
                id, driver := params["todriver"]
                if driver {
                    fmt.Println("user is sending a request to ", id)
                    strid, ok := id.(string)
                    if ok {
                        message := "user wants: " + strid
                        conn.WriteMessage(websocket.TextMessage, []byte(message))
                    }
                    //email messaging?
                    //private socket, private message between user-driver
                    //(if driver accepts)
                } else {
                    fmt.Println("theres not a driver to request")
                }
            }


        case websocket.BinaryMessage:
            fmt.Println(p)
        }

        err = conn.WriteMessage(websocket.TextMessage, []byte("received by server"))
        if err != nil {
            log.Println("error sending message")
            break
        }
    }

}


