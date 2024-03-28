package main

import ( 
    "fmt"
    "time"
    "math/rand"
    "net/http"
    "log"
    "os"

    "github.com/joho/godotenv"

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

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-type")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {
    //driver1:=new(Driver)
    //driver1.setLocation(1,1)
    //driver1.getCoordinate()

    err := godotenv.Load()
    if err != nil {
        fmt.Printf("error loading the environment variables")
    }
    myport := os.Getenv("PORT")
    environment := os.Getenv("ENVIRONMENT")
    port := ":" + myport

    if environment != "production" {
        fmt.Printf("in development\n")
    }

    mux := http.NewServeMux()
    handler := corsMiddleware(mux)

    mux.HandleFunc("/play", routes.Playthevideo)
    mux.HandleFunc("/check", routes.CheckProcess)

    mux.HandleFunc("/ws", socket.TheWSconn)

    mux.HandleFunc("/lookdriver", routes.LookforDrivers)//users looking for drivers
    mux.HandleFunc("/bevisible", routes.BeVisible) //for drivers


    mux.HandleFunc("/login", routes.Login)
    mux.HandleFunc("/register", routes.Register)
    mux.HandleFunc("/logout", routes.Logout)

    log.Printf("starting on localhost%s", port)
    http.ListenAndServe(port, handler)
}
