package routes

import (
    "log"
    "fmt"

    "net/http"
    "personal/wsservice/subprocess"
)

func CheckProcess(w http.ResponseWriter, r *http.Request){
    subprocess.ChecktheRunning("ffmpeg")
}

func Playthevideo (w http.ResponseWriter, r *http.Request) {
    log.Printf("playing the video")
    subprocess.StartFfmpeg()
    w.Header().Set("Contet-Type", "Text/plain")
    fmt.Fprintf(w, "playing the video")

}

func BeVisible(w http.ResponseWriter, r *http.Request){
    fmt.Printf("making the driver visible")
    //driver will send messages to all the users that near
    //I must take the real time dirver's location 
    //I must send location to user within the minimum area


    //redirect the dirver to a frontend page
    //this frontend page should start a ws connections
    //this frontend must sed location data
    http.Redirect(w,r, "ws://localhost:7777/ws", http.StatusTemporaryRedirect)
    fmt.Fprintf(w, "hello from home\n")
}

func LookforDrivers(w http.ResponseWriter, r *http.Request){
    log.Printf("%s /drivers\n",r.Method)
    w.Header().Set("Content-Type","text/plain")
    fmt.Fprintf(w, "drivers")
}

func Homesite(w http.ResponseWriter, r *http.Request){
    // I must take user's data
    log.Printf("%s /\n",r.Method)
    w.Header().Set("Content-Type","text/plain")
    fmt.Fprintf(w, "hello from home\n")
}

func WsHandler(w http.ResponseWriter, r *http.Request){
    log.Println("%s /ws", r.Method)
}
