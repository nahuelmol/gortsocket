package auth 

import (
    "fmt"

    "personal/wsservice/db"
)

func DoesuserExists(uname string, pass string) bool {
    rawcmd := "match username="+uname+" password="+pass
    result := db.PgQueries(rawcmd) //match command -> return true if user exists
    return result
}

//this file contains the login logic
//I must have access to the db here
func Login(username, password string) string {
    exists := DoesuserExists(username, password)
    if !exists {
        fmt.Println("the user doesn't exsits, first register\n")
        return "nil"
    }
    
    access_key, err := GenerateAccessKey()
    if err != nil {
        fmt.Println("err: ", err)
    }
    fmt.Println("access_key -> :", access_key)
    return access_key
}
