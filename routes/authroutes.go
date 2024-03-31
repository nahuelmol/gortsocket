package routes 

import (
    "net/http"
    "fmt"
    "personal/wsservice/authapi"
)
//this file just distributes the login and register logics by its routes
func Login(w http.ResponseWriter, r *http.Request){
    key := auth.Login()
    w.Header().Set("X-Access-Key", key)
    fmt.Fprintf(w, "logged in\n")
}

func Register(r http.ResponseWriter, w *http.Request){
    auth.Register()
}

func Logout(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "logged out\n")
    //I can do the logic of the logout right here
}
