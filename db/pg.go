package db

import (
    "fmt"
    "log"
    "strings"
    "database/sql"
    "errors"
    
    //"github.com/mattn/go-sqlite3"
)
//this is more sqlite3 than postgres actually

func pgConnection() *sql.DB {
    
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        fmt.Println(err)
    }

    return db
}

func UserDataout(entire string) map[string]string {
    var userData = make(map[string]string)
    exists_kv := false
    var (
        key string
        value string
    )

    words := strings.Split(entire, " ")
    for _, word := range words {
        for _, char := range word {
            if char == '=' {
                exists_kv = true
                key = value
                value = ""
            } else {
                value += string(char)
            }
        }
        if exists_kv {
            userData[key] = value
            key = ""
            value = ""
        } else {
            key = ""
            value = ""
        }
    }
    return userData
}

func ActionOut(entire string) (string, error) {
    var action string
    for _, c := range entire {
        if c == ' ' {
            return action, nil
        }
        action += string(c)
    }
    return "nil", errors.New("there's not an action")
}

func matchDB(username, password string) bool {
    db := pgConnection()
    sqlcmd := fmt.Sprintf("SELECT %s FROM users;", username)
    rows, err := db.Query(sqlcmd)
    if err != nil {
        fmt.Printf("error: %v", err)
    }
    //I need to get db password and compare it to the recently user's password

    var values []interface{}
    var all string
    for rows.Next() {
        if err := rows.Scan(values...); err != nil {
            log.Fatal(err)
        }
    }

    for _, val := range values {
        all += fmt.Sprintln(val)
    }

    fmt.Printf("password:", password)
    fmt.Println("passowrd in db:", all)

    if all == password {
        db.Close()
        return true
    } else {
        db.Close()
        return false
    }
    db.Close()
    return false
}

func delUser(username string){
    db := pgConnection()
    rawcmd := fmt.Sprintf("INSERT INTO users(name) VALUES(%s)", username)
    val, err := db.Exec(rawcmd)
    if err != nil {
        fmt.Printf("error: %v", err)
    }
    fmt.Println(val)
    db.Close()
}


func addUser(username, password string){
    db := pgConnection()
    rawcmd := fmt.Sprintf("DELETE FROM users WHERE name = %s", username)
    val, err := db.Exec(rawcmd)
    if err != nil {
        fmt.Printf("error: %v", err)
    }
    fmt.Println(val)
    db.Close()
}

func PgQueries(cmd string) bool {
    //this works with "match username=nahuel password=verstappen33"
    action, err := ActionOut(cmd)
    user        := UserDataout(cmd)
    username := user["username"]
    password := user["password"]

    if err != nil {
        fmt.Println("error handling the action")
        return false
    }

    switch action {
    case "add":
        addUser(username, password)
        fmt.Printf("adding")
    case "del":
        fmt.Printf("deleting")
    case "upd":
        fmt.Printf("updating")
    case "match":
        doesmatch := matchDB(username, password)
        if doesmatch {
            fmt.Printf("user account exists")
            return true 
        } else {
            return false
        }

    default:
        fmt.Printf("unrecognized command")
    }

    return false
}
