package main

import (
    //"net/http"
    "github.com/microcats/hellgate/proxy"
)

func main() {
    proxy.NewMultipleHostReverseProxy()
    //http.Handle("/", r)
    //http.ListenAndServe(":8000", r)
}
