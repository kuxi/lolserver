package main

import (
    "io"
    "net/http"
    "log"
    "fmt"
    "controllers"
)

func Serve(c controllers.Controller) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, req *http.Request) {
        response := c.Handler(req)
        if response.Status == http.StatusOK {
            io.WriteString(w, response.Body)
        } else if response.Status == http.StatusMovedPermanently {
            http.Redirect(w, req, response.Body, response.Status)
        } else {
            http.Error(w, response.Body, response.Status)
            fmt.Printf("Response not 200")
        }
    }
}

func main() {
    controllers := controllers.GetControllers()
    for _,c := range controllers {
        http.HandleFunc("/hello", Serve(c))
    }
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
