package routes 

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"

    "personal/wsservice/authapi"
)
//this file just distributes the login and register logics by its routes
func Login(w http.ResponseWriter, r *http.Request){
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }
    var bodymap map[string]interface{}

    err = json.Unmarshal(body, &bodymap)
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    username := bodymap["username"].(string)
    password := bodymap["password"].(string)

    key := auth.Login(username, password)
    w.Header().Set("X-Access-Key", key)
    fmt.Fprintf(w, "logged in\n")
}

func Register(w http.ResponseWriter, r *http.Request){
    access_key := r.Header.Get("X-Access-Key")
    if len(access_key) > 0  {
        fmt.Fprintf(w, "you are already logged in\n")
        return
    }
    
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }
    var bodymap map[string]interface{}
    err = json.Unmarshal(body, &bodymap)
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    username := bodymap["username"].(string)
    password := bodymap["password"].(string)

    auth.Register(username, password)
    fmt.Fprintf(w, "registered in\n")
}

func Logout(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "logged out\n")
    //I can do the logic of the logout right here
}
