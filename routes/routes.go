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
