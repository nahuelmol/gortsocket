package auth

import (
    "fmt"

    "personal/wsservice/db"
)

func Register(username, password string){
    //I must have access the database here
    exists := DoesuserExists(username, password)
    if !exists {
        fmt.Println("the user doesn't exists, first register\n")
    }   
    cmd := fmt.Sprintf("add username=%s password=%s", username, password)
    db.PgQueries(cmd)
    //generating access key
    access_key, err := GenerateAccessKey()
    if err != nil {
        fmt.Println("err: ", err)
    }
    fmt.Println("access_key -> :", access_key)
}
