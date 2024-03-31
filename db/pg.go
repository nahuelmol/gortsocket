package db

import (
    "fmt"
    "strings"
    "database/sql"
    
    "github.com/lib/pq"
)


func pgConnection() sql.Open {
    user := "nahuel"
    dbname := "accountsdb"
    sslmode := "verify-full"

    connStr := user+ " " +dbname+ " " +sslmode 
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    return db
}

func UserDataout(entire string) string {
    var userData = make(map[string]string)
    exists_kv := false
    var (
        key string
        value string
    )

    words := strings.Split(entire, " ")
    for _, word := range words {
        for i, char := range word {
            if char == "=" {
                exists_kv = true
                key = value
                value = ""
            } else {
                value += char
            }
        }
        if exists_kv {
            userData[key] = term
            key = ""
            term = ""
        }
    }
}

func ActionOut(entire string) string {
    var action string
    for _, c := range entire {
        if c == ' ' {
            return action
        }
        action += c
    }
    return nil
}

func matchDB(){
    db := pgConnection()
    sqlcmd := fmt.Sprintf("SELECT * FROM users;")
    row, err := db.Query(sqlcmd)
}

func addUser(username, password string){
    rawcmd := fmt.Sprintf("INSERT INTO users(name)
	VALUES(%s) RETURNING id", username)
    err := db.QueryRow(rawcmd).Scan(&userid)
}

func PgQueries(cmd string) {
    action := ActionOut(cmd)

    switch action {
    case 'add':
        addUser(username, password)
        fmt.Printf("adding")
    case 'del':
        fmt.Printf("deleting")
    case 'upd':
        fmt.Printf("updating")
    case 'match':
        matchDB(username, password)
        fmt.Printf("matching")

    default:
        fmt.Printf("unrecognized command")
}

