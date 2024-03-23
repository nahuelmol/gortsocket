
package db

import (
    "log"
    "github.com/redis/go-redis/v9"
)
type RedisClient struct {
    Addr
    Password
    DB
}

func CreateClient (Addr, Password, db string) *RedisClient {
    return &RedisClient {
        Addr,
        Password,
        db,
    }
}

func conn() *redis.NewClient {
    client_object := CreateClient("","","")
    rclient := redis.NewClient(&redis.Options{RedisClient})

    return rclient
}

