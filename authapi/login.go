package auth 

import (
    "fmt"
    "personal/wsservice/utils"
)

//this file contains the login logic
//I must have access to the db here
func Login(){
    access_key, err := GenerateAccessKey()
    fmt.Println("access_key -> :", access_key)
}
