package main

import (
    "io"
    "io/ioutil"
    "bytes"
    "net/http"
    "log"
    "github.com/hoisie/mustache"
    "fmt"
)

func ReadAll(r io.Reader) string {
    buf := new(bytes.Buffer)
    buf.ReadFrom(r)
    return buf.String()
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
    var context = map[string]interface{}{}
    str := mustache.RenderFile("templates/hello.html", context)
    if req.Method == "POST" {
        file, _, e := req.FormFile("file")
        if e != nil {
            fmt.Printf("Error:%s\n", e)
        }
        buf := new(bytes.Buffer)
        buf.ReadFrom(file)
        e = ioutil.WriteFile("test.jpg", buf.Bytes(), 0660)
        if e != nil {
            fmt.Printf("Error writing:%s\n", e)
        }
    }
    io.WriteString(w, str)
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
