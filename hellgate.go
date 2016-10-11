package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "time"
    "fmt"
    "github.com/gorilla/mux"
)

//Target url: http://apache.org/server-status
//Url through proxy:  http://localhost:3002/forward/server-status


func main() {
    target := "http://127.0.0.1:9090"
    remote, err := url.Parse(target)
    if err != nil {
        panic(err)
    }

    proxy := httputil.NewSingleHostReverseProxy(remote)
    r := mux.NewRouter()
    r.HandleFunc("/test/{rest:.*}", handler(proxy))
    http.Handle("/", r)
    http.ListenAndServe(":8000", r)
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = mux.Vars(r)["rest"]
        t := time.Now()
        p.ServeHTTP(w, r)
        fmt.Println(time.Since(t).Seconds())
    }
}
