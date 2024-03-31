package socket

import (
    "fmt"
    //"encoding/json"
)

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
