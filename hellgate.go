package main

import(
    "log"
    "net/http"
    "github.com/microcats/hellgate/proxy"
)

func main() {
    proxy.New("test", "http://127.0.0.1:9090/").Register()
    listen()
}

func listen() {
    log.Fatal(http.ListenAndServe(":8080", nil))
}
