package db

import (
    "github.com/redis/go-redis/v9"
)

type RedisClient struct {
    Addr string
    Password string
    DB int
    Protocol int
}

func SetOptions()  *redis.Options {
    return &redis.Options{
        Addr: "",
        Password: "",
        DB: 0,
        Protocol:3,
    }
}

func Conn() *redis.Client {
    client := SetOptions()
    rclient := redis.NewClient(client)

    return rclient
}


