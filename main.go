package main

import (
    "fmt"
    "net"
    "net/http"
    "net/http/fcgi"
)

func handler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Hello World!")
}

func main() {
    l, err := net.Listen("tcp", "127.0.0.1:9000") // TCP 9000 番ポートで Listen
    if err != nil {
        return
    }
    http.HandleFunc("/", handler)
    fcgi.Serve(l, nil)
}