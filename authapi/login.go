package auth 

import (
    "fmt"
)

//this file contains the login logic
//I must have access to the db here
func Login(){
    access_key, err := GenerateAccessKey()
    if err != nil {
        fmt.Println("err: ", err)
    }
    fmt.Println("access_key -> :", access_key)
}
