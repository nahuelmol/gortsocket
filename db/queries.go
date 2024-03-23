package db

import (
    "log"
    "personal/wsservice/db"
    "github.com/go-redis/v9"
)

func SetPair(key, value string){
    conn := db.Conn()
    err := conn.Set(ctx, key, value, 0).Err()
    if err != nil { 
        log.Println("error setting")
    }
    
    err := conn.Get(ctx, key, value, 0).Err()
    if err != nil {
        log.Println("error getting")
    }
}
func GetPair(){
    conn := db.Conn()
    err := conn.Set(ctx, key, value, 0).Err()
    if err != nil { 
        log.Println("error setting")
    }
    
    err := conn.Get(ctx, key, value, 0).Err()
    if err != nil {
        log.Println("error getting")
    }
}
