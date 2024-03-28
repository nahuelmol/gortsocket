package routes 

import (
    "net/http"
    
    "personal/wsservice/authapi"
)
//this file just distributes the login and register logics by its routes
func Login(r http.ResponseWriter, w *http.Request){
    auth.Login()
}

func Register(r http.ResponseWriter, w *http.Request){
    auth.Register()
}

func Logout(r http.ResponseWriter, w *http.Request){
    //I can do the logic of the logout right here
}
