package main

import ( 
    "fmt"
    "time"
    "math/rand"
    "net/http"
    "log"

    "personal/wsservice/routes"
    "personal/wsservice/wsocket"
)

type IPv4 int
type IPv6 int

type Coordinate struct {
    xposition int32 
    yposition int32

    pTime time.Time
}

type User struct {
    coordiante Coordinate
    service string //type of service the user looks for
    id_sector int32
    drivers_ammount int32 //drivers available in user's sector
    drivers[]int32 //this should be an slice, a mutable collection
}

type Driver struct {
    driver_id int32
    coor Coordinate //driver's coordinates
}

func (driver *Driver) setLocation(x,y int32) {
    //function that takes the current time and save it into a time variable
    var mycoor Coordinate
    mycoor.xposition = x
    mycoor.yposition = y
    mycoor.pTime = time.Now()

    driver.driver_id = int32(rand.Intn(101))
    driver.coor = mycoor
}
func (driver *Driver) getCoordinate() {
    fmt.Printf("driver id:%d\n", driver.driver_id)
    fmt.Printf("driver x:%d \n", driver.coor.xposition) 
    fmt.Printf("driver y:%d \n" , driver.coor.yposition) 
    fmt.Printf("driver time:%d\n", driver.coor.pTime.Hour()) 
}

func main() {
    //driver1:=new(Driver)
    //driver1.setLocation(1,1)
    //driver1.getCoordinate()
    http.HandleFunc("/drivers", routes.LookforDrivers)
    http.HandleFunc("/", routes.Homesite)
    http.HandleFunc("/wshandler", routes.WsHandler)

    http.HandleFunc("/play", routes.Playthevideo)
    http.HandleFunc("/check", routes.CheckProcess)

    http.HandleFunc("/ws", socket.TheWSconn)

    log.Printf("starting a basic server on port 8080")
    http.ListenAndServe(":8080", nil)
}
