package auth 

import (
    "fmt"
)

func DoesuserExists(uname string, pass string) bool {
    rawcmd := "match username="+uname+" password="+pass
    result := PgQueries(rawcmd)
    return false
}

//this file contains the login logic
//I must have access to the db here
func Login() string {
    exists := DoesuserExists(username, passoword)
    if !exists {
        fmt.Println("the user doesn't exsits, first register\n")
        return nil
    }
    
    access_key, err := GenerateAccessKey()
    if err != nil {
        fmt.Println("err: ", err)
    }
    fmt.Println("access_key -> :", access_key)
    return access_key
}
