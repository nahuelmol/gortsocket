package auth

import (
    "fmt"
)

func Register(){
    //I must have access the database here
    access_key, err := GenerateAccessKey()
    if err != nil {
        fmt.Println("err: ", err)
    }
    fmt.Println("access_key -> :", access_key)
}
