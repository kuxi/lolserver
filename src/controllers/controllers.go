package controllers
import (
    "bytes"
    "io/ioutil"
    "net/http"
    "github.com/hoisie/mustache"
    "fmt"
)

type Response struct {
    Status int
    Body   string
}

func ResponseOK(body string) Response {
    return Response {
        Status: http.StatusOK,
        Body: body,
    }
}

func ResponseTeapot(body string) Response {
    return Response {
        Status: http.StatusTeapot,
        Body: body,
    }
}

func ResponseBadRequest(body string) Response {
    return Response {
        Status: http.StatusBadRequest,
        Body: body,
    }
}

func ResponseServerError(body string) Response {
    return Response {
        Status: http.StatusInternalServerError,
        Body: body,
    }
}

func ResponseRedirect(url string) Response {
    return Response {
        Status: http.StatusMovedPermanently,
        Body: url,
    }
}

type Controller struct {
    Pattern string
    Handler func(req *http.Request) Response
}

func GetControllers() []Controller {
    uploadController := Controller {
        Pattern: "/hello",
        Handler: UploadHandler,
    }
    return []Controller {
        uploadController,
    }
}

func UploadHandler(req *http.Request) Response {
    var context = map[string]interface{}{}
    str := mustache.RenderFile("templates/hello.html", context)
    if req.Method == "POST" {
        file, _, e := req.FormFile("file")
        if e != nil {
            fmt.Printf("Error:%s\n", e)
            return ResponseBadRequest("file parameter missing")
        }
        buf := new(bytes.Buffer)
        buf.ReadFrom(file)
        e = ioutil.WriteFile("test.jpg", buf.Bytes(), 0660)
        if e != nil {
            fmt.Printf("Error writing:%s\n", e)
            return ResponseServerError("Unable to save file")
        }
        return ResponseRedirect("/hello")
    }
    return ResponseOK(str)
}
