package main

import (
    "fmt"
    "time"
    "runtime"
    "net/http"
    "github.com/microcats/hellgate/proxy"
    "github.com/microcats/hellgate/backend"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    t := time.Now()

    machines := []string{"http://127.0.0.1:2379"}
    storeClient, err := backend.NewEtcdClient(machines, "", "", "", false, "", "")
    if err != nil {
        panic(err)
    }

    r, err := proxy.NewMultipleHostReverseProxy(storeClient)
    if err != nil {
        panic(err)
    }

    fmt.Println(time.Since(t).Seconds())
    http.Handle("/", r)
    http.ListenAndServe(":8000", r)
}
