package socket

import (
    "fmt"
    "errors"

    "personal/wsservice/obj"
    "personal/wsservice/distance"
)

func toStackLocationUser(action string, node *obj.Node) func() bool {
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

func toStackLocationDriver(action string, node *obj.Node) func() bool {
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

func toStackDistance(action string, node *obj.Node) func() bool {
    stack := distance.CreateStack()
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
    driver  := new(obj.Driver)
    coor    := driver.SetLocation(int32(xloc), int32(yloc))
    node    := obj.CreateNode(coor)

    toStackLocationDriver("push", node)//driver
    toStackLocationUser("push", node) //user

    lastDriver  := toStackLocationDriver("top")
    lastUser    := toStackLocationUser("top")
    distance    := CalculateDistance(lastDriver, lastUser)

    toStackDistance("psuh", distance)
}

func BuildgroupMsg(id uint32, typeof string) string {
    userStack := UserRegister[id]

    data, err := userStack.Topdata()
    if err != nil {
        fmt.Println(err)
        return "nil data"
    }

    fmt.Println(data)
    //position := make(map[string]string)
    //position["x"] = string(data.xposition)
    //position["y"] = string(data.yposition)
    //position["t"] = string(data.pTime)

    //json_str, err := json.Marshal(position)
    //if err != nil {
    //    fmt.Println("there's a conversion error")
    //}
    

    //message := json_str
    //return string(message)
    return "hello"
}
type Message struct {
    id uint32
    name string
}

type SessionMessage struct {
    Message
    access_key string
}

type RawMessage struct {
    content string
}

func (rm RawMessage) String() string {
    return rm.content
}


func CreateMsg(params map[string]interface{}) {
    username, isuser := params["username"].(string)
    if !isuser {
        username = "random"
    }
    access_key, ak := params["access_key"].(string)

    id, isid := params["id"].(uint32)
    if !isid {
        id = 0
    }

    if ak {
        session := SessionMessage { 
            Message { name:username, id:id,},
            access_key,
        }
        fmt.Println(session)
        //StackSession(session)
    } else {
        msg := Message {name:username, id:params["id"].(uint32)}
        fmt.Println(msg)
    }

    //message stacking
    //just message with an id are stacked
}

func StartSessionMsg(params map[string]interface{}) (bool, error){
     
    //this will take those messages that start sessions
    id, is_id := params["id"].(uint32)
    if is_id {
        CreateMsg(params)
        //db.AccessKeyInDB(access_key, id)
        xloc, xok := params["xlocation"].(uint32)
        yloc, yok := params["ylocation"].(uint32) 
        if !xok || !yok {
            return false, errors.New("error taking location data")
        }

        locate(id, xloc, yloc)
        return true, nil
    }
    return false, errors.New("not exists an id") 
}

func LocationMsg() {}

func Rawmessage(content string) string {
    return content
}
