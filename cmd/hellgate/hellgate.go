package main

import (
    "fmt"
    "time"
    "net/http"
    "github.com/microcats/hellgate/proxy"
)

func main() {
    t := time.Now()
    r, err := proxy.NewMultipleHostReverseProxy()
    if err != nil {
        panic(err)
    }
    fmt.Println(time.Since(t).Seconds())
    http.Handle("/", r)
    http.ListenAndServe(":8000", r)
}
