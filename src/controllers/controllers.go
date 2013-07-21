package controllers

type Response struct {
    Status int
    Body   string
}

func ResponseOK(body string) Response {
    return Response {
        Status = 200
        Body = body
    }
}

type Controller struct {
    Path    string
    Handler func(req *http.Request) Response
}
