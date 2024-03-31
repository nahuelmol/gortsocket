package main

import ( 
    "fmt"
    "net/http"
    "log"
    "os"

    "github.com/joho/godotenv"

    "personal/wsservice/obj"
    "personal/wsservice/routes"
    "personal/wsservice/wsocket"
)

type IPv4 int
type IPv6 int

var hmanyConnections int = 0

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

func exampleWare(next http.Handler) http.Handler {
    nh := func(w http.ResponseWriter, r *http.Request){

        fmt.Fprintf(w, "starting\n")
        next.ServeHTTP(w, r)
        fmt.Fprintf(w, "\nfinished")
    }
    return http.HandlerFunc(nh)
}

func main() {
    //driver1:=new(Driver)
    //driver1.setLocation(1,1)
    //driver1.getCoordinate()
    UserRegister := make(map[uint32]*obj.StackLocation)
    fmt.Println(UserRegister)

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

    mux     := http.NewServeMux()
    handler := corsMiddleware(mux)

    mux.HandleFunc("/play", routes.Playthevideo)
    mux.HandleFunc("/check", routes.CheckProcess)

    mux.Handle("/ws", exampleWare(http.HandlerFunc(socket.TheWSconn)))

    mux.Handle("/lookdriver", exampleWare(http.HandlerFunc(routes.LookforDrivers)))//users looking for drivers
    mux.HandleFunc("/bevisible", routes.BeVisible) //for drivers

    mux.HandleFunc("/login", routes.Login)
    mux.HandleFunc("/register", routes.Register)
    mux.HandleFunc("/logout", routes.Logout)

    log.Printf("starting on localhost%s", port)
    http.ListenAndServe(port, handler)
}
