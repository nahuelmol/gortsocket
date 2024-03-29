package db

import (
    "context"
    "log"
    "time"
    "fmt"
)

var ctx = context.Background()

func SetPair(key, value string){
    conn := Conn()
    err := conn.Set(ctx, key, value, 0).Err()
    if err != nil { 
        log.Println("error setting")
    }
}

func GetPair(key string){
    conn := Conn()
    
    val, err := conn.Get(ctx, key).Result()
    if err != nil {
        log.Println("error getting")
    } else {
        log.Printf("the value is: %s", val)
    }
}

func AccessSetting(access_key, user string) bool {
    lifetime := 2 * time.Hour

    conn := Conn()
    err := conn.Set(ctx, access_key, user, lifetime).Err()
    if err != nil {
        log.Println("error setting access key")
        return false
    } else {
        log.Println("value correctly setted")
        return true
    }
}

func AccessKeyInDB(access_key, user string) bool {
    fmt.Println("checking if access key for the user exists in db")
    conn := Conn()
    
    val, err := conn.Get(ctx, access_key).Result()
    if err != nil {
        log.Println("error getting")
        return false
    } else {
        if val == user {
            log.Printf("value: %s exists", val)
            return true
        } else {
            //the access key exists but its owner is not user that is requireing it right now
            return false
        }
    }

}
