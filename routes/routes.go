package routes

import (
    "log"
    "fmt"
    "sync"
    "time"
    "strconv"

    "net/http"
    "personal/wsservice/obj"
    "personal/wsservice/wsocket"

    "github.com/gorilla/websocket"
)
var Mutex sync.Mutex

func streamLocations(ch chan string){
    for i := 1; i <= 666; i++ {
        istring := strconv.Itoa(i)
        ch <- istring
        time.Sleep(1 * time.Second)
    }
    close(ch)
}

func BeVisible(w http.ResponseWriter, r *http.Request){
    fmt.Printf("making the driver visible")
    ch := make(chan string)
    go streamLocations(ch)
    //driver will send messages to all the users that are near
    //I must take the real time driver's location 
    //I must send location to user within the minimum area

    /*latStr := r.URL.Query().Get('lat')
    lngStr := r.URL.Query().Get('lng')
    lat, err := strconv.ParseFloat(latStr, 64)
    if err != nil {
        http.Error(w, "Invalid latitude value")
        return
    }
    lng, err := strconv.ParseFlow(lngStr, 64)
    if err != nil {
        http.Error(w, "Invalie longitude value")
        return
    }*/

    driver := new(obj.Driver)
    Slocation := obj.CreateStack()
    
    coor := driver.SetLocation(4,4) //random initial location
    node := obj.CreateNode(coor)
    
    driver_id := driver.Identifier()
    socket.UserRegister[driver_id] = Slocation
    //UserRegister[driver_id] = Slocation

    // I am going to stack nodes (coordinate nodes)
    Slocation.Push(node)
    //redirect the driver to a frontend page
    //this frontend page should start a ws connections
    //this frontend must sed location data
    fmt.Println("\nnode stacked")
    message := "\ndriver was activated"

    broadcast := []byte("hello,clients!")
    Mutex.Lock()
    for {
        data, ok := <-ch;
        if !ok {
            break;
        }
        broadcast = []byte(data)
        for client := range socket.Clients {
            if err := client.WriteMessage(websocket.TextMessage, broadcast); err != nil {
                client.Close()
                delete(socket.Clients, client)
            }
        }
    }
    Mutex.Unlock()
    fmt.Fprintf(w, message)
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
