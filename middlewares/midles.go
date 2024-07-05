package middlewares

import (
    "net/http"
    "strings"
    "log"
)

func CookiesMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        cookies := r.Cookies()
        for _, cookie := range cookies {
            log.Printf("cookie name: ", cookie.Name)
        }
    })
}

func LogMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w,r)
    })
}

func JSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        if strings.HasSuffix(r.URL.Path, ".js"){
            w.Header().Set("Content-Type", "application/javascript")
        }
        next.ServeHTTP(w,r)
    })
}

func CorsMiddleware(next http.Handler) http.Handler {
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
