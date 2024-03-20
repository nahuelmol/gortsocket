package routes

fmt (
    "log"
    "fmt"

    "gtihub.com/gorilla/socket"
)

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

func Wshandler(w http.ResponseWriter, r *http.Request){
    log.Println("%s /ws", r.Method)
    socket.WsocketHandler(w,r)
}
