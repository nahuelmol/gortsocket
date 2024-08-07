package routes

import (
    "log"
    "fmt"
    "sync"
    "time"
    "strconv"
    "math/rand"
    "net/http"
    "html/template"

    //"personal/wsservice/obj"
    "personal/wsservice/wsocket"

    "github.com/gorilla/websocket"
)
var Mutex sync.Mutex

func streamLocations(ch chan string){
    for i := 1; i <= 666; i++ {
        rand.Seed(time.Now().UnixNano())
        randomInt := rand.Intn(100)
        istring := strconv.Itoa(randomInt)
        ch <- istring
        time.Sleep(1 * time.Second)
    }
    close(ch)
}

func StreamLocation(ch chan string, location string){
        ch <- location
}

//locationCH := make(chan string,1024)
//go streamLATLON(ch2)

var locationCH chan string
func BeVisible(w http.ResponseWriter, r *http.Request){
    fmt.Printf("making the driver visible")
    ch := make(chan string,3)
    

    latStr := r.URL.Query().Get("lat")
    lonStr := r.URL.Query().Get("lon")
    lat, err := strconv.ParseFloat(latStr, 64)
    if err != nil {
        http.Error(w, "Invalid latitude value", http.StatusNotFound)
        return
    }
    lon, err := strconv.ParseFloat(lonStr, 64)
    if err != nil {
        http.Error(w, "Invalie longitude value", http.StatusNotFound)
        return
    }
    latRes := strconv.FormatFloat(lat, 'f', 4, 64)
    lonRes := strconv.FormatFloat(lon, 'f', 4, 64)
    location := "["+latRes +","+ lonRes+"]"

    Mutex.Lock()
    go StreamLocation(ch, location)
    Mutex.Unlock()
    //driver := new(obj.Driver)
    //Slocation := obj.CreateStack()
    
    //coor := driver.SetLocation(4,4) //random initial location
    //node := obj.CreateNode(coor)
    
    //driver_id := driver.Identifier()
    //socket.UserRegister[driver_id] = Slocation
    //UserRegister[driver_id] = Slocation

    // I am going to stack nodes (coordinate nodes)
    //Slocation.Push(node)
    //redirect the driver to a frontend page
    //this frontend page should start a ws connections
    //this frontend must sed location data
    fmt.Println("\nnode stacked")

    Mutex.Lock()
    go onceWriter(ch)
    Mutex.Unlock()
}

func onceWriter(ch chan string){
    data, ok := <-ch;
    fmt.Println("ch:", data)
    if !ok {
        return
    }
    broadcast := []byte(data)
    for client := range socket.Clients {
        if err := client.WriteMessage(websocket.TextMessage, broadcast); err != nil {
            client.Close()
            delete(socket.Clients, client)
        }
    }
}
func clientWriter(ch chan string){
    for {
        data, ok := <-ch;
        if !ok {
            break;
        }
        broadcast := []byte(data)
        for client := range socket.Clients {
            if err := client.WriteMessage(websocket.TextMessage, broadcast); err != nil {
                client.Close()
                delete(socket.Clients, client)
            }
        }
    }
}

func LookforDrivers(w http.ResponseWriter, r *http.Request){
    log.Printf("%s /drivers\n",r.Method)
    w.Header().Set("Content-Type","text/plain")
    fmt.Fprintf(w, "drivers")
}

func Home(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("public/views/index.html"))
    err  := tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func WsHandler(w http.ResponseWriter, r *http.Request){
    log.Println("%s /ws", r.Method)
}
