package socket

import (
    "fmt"
    "errors"
    "reflect"
    "time"

    "personal/wsservice/obj"
)

func toStackLocationUser() func(string, *obj.Node) (*obj.Node, error){
    stack := obj.CreateStack()
    return func(action string, node *obj.Node) (*obj.Node, error) {
        switch action {
        case "push":
            stack.Push(node)
            return nil, nil
        case "pop":
            stack.Pop()
            return nil, nil
        case "top":
            result, err := stack.Topdata()
            if err != nil {
                return nil, errors.New("just error")
            }
            return result, nil
        case "next":
            stack.Nextdata()
            return nil, nil
        case "whole":
            stack.Wholedata()
            return nil, nil
        case "lenght":
            stack.Getlength()
            return nil, nil
        default:
            fmt.Println("unrecognized action")
            return nil, nil
        }
    }
}

func toStackLocationDriver() func(action string, node *obj.Node) (*obj.Node, error) {
    stack := obj.CreateStack()
    return func(action string, node *obj.Node) (*obj.Node, error) {
        switch action {
        case "push":
            stack.Push(node)
            return nil, nil
        case "pop":
            stack.Pop()
            return nil, nil
        case "top":
            result, err := stack.Topdata()
            if err != nil {
                return nil, errors.New("just error")
            }
            return result, nil
        case "next":
            stack.Nextdata()
            return nil, nil
        case "whole":
            stack.Wholedata()
            return nil, nil
        case "lenght":
            stack.Getlength()
            return nil, nil
        default:
            fmt.Println("unrecognized action")
            return nil, nil
        }
    }
}

func ToDistance() func(string, *obj.Distance) (*obj.Distance, error) {
    stack := obj.CreateDistancer()
    return func(action string, distance *obj.Distance) (*obj.Distance, error) {
        switch action {
        case "push":
            stack.Push(distance)
            return nil, nil
        case "pop":
            stack.Pop()
            return nil, nil
        case "top":
            lastDistance := stack.Gethead()
            return lastDistance, nil
        default:
            fmt.Println("unrecognized action")
            return nil, nil
        }
    }
}

func Takeout(result *obj.Node, _ error) obj.Coordinate {
    val := reflect.ValueOf(result)
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
    data := val.FieldByName("data")
    if data.IsValid() {
        fmt.Println("is valid")
    } else {
        fmt.Println("not valid")
    }

    var x int32
    var y int32
    xx := data.FieldByName("Xposition")
    yy := data.FieldByName("Yposition")

    if xx.Kind() == reflect.Int32 || yy.Kind() == reflect.Int32 {
        xvalue := int32(xx.Int())
        yvalue := int32(yy.Int())
        x = xvalue
        y = yvalue
    } else {
        fmt.Println("is not an int32")
    }

    rightnow := time.Now()
    coor := obj.Coordinate { Xposition:x, Yposition:y, Time:rightnow }
    return coor 
}

func locate(id, xloc, yloc uint32) {
    driver  := new(obj.Driver)
    coor    := driver.SetLocation(int32(xloc), int32(yloc))
    node    := obj.CreateNode(coor)

    stack_driv := toStackLocationDriver()//driver
    stack_user := toStackLocationUser() //user

    stack_driv("push", node)

    lastdriver  := Takeout(stack_driv("top", nil))
    lastuser    := Takeout(stack_user("top", nil))

    distance    := obj.CalculateDistance(lastdriver, lastuser)
    dist_instance := obj.CreateDistance(float64(distance))
    
    distancer := ToDistance()
    distancer("push", dist_instance)

    head, err := distancer("top", nil)
    if err != nil {
        fmt.Println("error taking the head")
    }

    HEAD := reflect.ValueOf(head)
    if HEAD.Kind() == reflect.Ptr {
        HEAD = HEAD.Elem()
    }
    val := HEAD.FieldByName("val")

    var topdistance float64
    if val.Kind() == reflect.Float64 {
        topdistance = float64(val.Float())
    } else {
        fmt.Println("the distance is not even a float64")
    }

    fmt.Println("topdistance:", topdistance)
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
