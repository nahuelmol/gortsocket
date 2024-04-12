package socket 

import ( 
    "log"
    "net/http"
    "encoding/json"
    "sync/atomic"
    "fmt"
    
    "personal/wsservice/db"
    "personal/wsservice/obj"
    
    "github.com/gorilla/websocket"
)

var HmanyConnections int32
var UserRegister map[uint32]*obj.StackLocation

func toStack(action string, node *obj.Node) func() bool {
    stack := obj.CreateStack()
    return func() bool {
        switch action {
        case "push":
            stack.Push(node)
        case "pop":
            stack.Pop()
        case "top":
            stack.Topdata()
        case "next":
            stack.Nextdata()
        case "whole":
            stack.Wholedata()
        case "lenght":
            stack.Getlength()
        default:
            fmt.Println("unrecognized action")
        }
        return true
    }
}


func locate(id, xloc, yloc uint32) {
    driver := new(obj.Driver)
    coor := driver.SetLocation(int32(xloc), int32(yloc))
    node := obj.CreateNode(coor)

    toStack("push", node)
}


var upgrader = websocket.Upgrader {
        ReadBufferSize: 1024,
        WriteBufferSize:1024,
}


func TheWSconn(w http.ResponseWriter, r *http.Request){
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading")
    }
    contextValue := r.Context().Value("v")
    fmt.Println("context value:", contextValue)

    defer conn.Close()
    for {
        msgType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("error reading message")
            log.Println(err)
            break
        }
        atomic.AddInt32(&HmanyConnections, 1)

        switch msgType {
        case websocket.TextMessage:
            json_string := string(p)
            var params map[string]interface{}
            json.Unmarshal([]byte(json_string), &params)

            source, source_exists := params["source"]
            if !source_exists {
                message := "source is nil"
                err = conn.WriteMessage(websocket.TextMessage, []byte(message))
            } else {
                fmt.Println("source: ", source)
            }

            username, uname_exists := params["username"]
            if uname_exists {
                message := "hi " + username.(string) 
                err = conn.WriteMessage(websocket.TextMessage, []byte(message))
            } else {
                message := "username is nil"
                err = conn.WriteMessage(websocket.TextMessage, []byte(message))
            }

            access_key, ak := params["access_key"].(string)
            did, okd := params["driverid"].(uint32) 
            uid, oku := params["userid"].(uint32) 

            if !ak {
                message := "not an access key"
                err = conn.WriteMessage(websocket.TextMessage, []byte(message))
                if err != nil {
                    fmt.Println("error writing the msg")
                }

                if uname_exists {
                    fmt.Println("some user trying to conn")
                }
                if okd {
                    fmt.Println("identified driver trying conn")
                } else { 
                    if oku {
                        fmt.Println("identified user triyng conn!")
                    } else {
                        message = "your are generic, bad day"
                        conn.WriteMessage(websocket.TextMessage, []byte(message))
                    }
                }
            } else {
                src := "unknown"
                var id uint32
                if oku && !okd {
                    src = "user"
                    id = uid
                    fmt.Println("it's a user")
                } else if okd && !oku {
                    src = "driver"
                    id = did
                    fmt.Println("it's a driver")

                    //put its location in the stack
                    xloc, xok := params["xlocation"].(uint32) 
                    yloc, yok := params["ylocation"].(uint32) 

                    if !xok || !yok {
                        return 
                    }

                    locate(id, xloc, yloc)
                } else if !oku && !okd {
                    fmt.Println("it's not a driver or user")
                } else {
                    src = "monster"
                }

                message := BuildgroupMsg(id, src)
                err = conn.WriteMessage(websocket.TextMessage, []byte(message))

                db.AccessKeyInDB(access_key, id)
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


