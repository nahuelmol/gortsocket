package auth

import (
    "fmt"
    "personal/wsservice/utils"
)

func Register(){
    //I must have access the database here
    access_key, err := GenerateAccessKey()
    fmt.Println("access_key -> :", access_key)
}
