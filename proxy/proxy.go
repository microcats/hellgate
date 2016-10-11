package proxy

import (
    //"log"
    "fmt"
    //"time"
    //"net/url"
    //"net/http"
    //"net/http/httputil"
    //"github.com/gorilla/mux"
    "github.com/microcats/hellgate/backend"
    //"github.com/coreos/etcd/client"
    //"context"
)


func NewMultipleHostReverseProxy() {
    machines := []string{"http://127.0.0.1:2379"}
    etcd, _ := backend.NewEtcdClient(machines, "", "", "", false, "", "")
    result, _ := etcd.Get("/hellgate/apis")
    //{"prefix":"test1","upstream_url":"http://127.0.0.1:9090", "create_at":"2016-02-01 15:11:22"}
    for _, value := range result.Node.Nodes {
        list, _ := etcd.Get(value.Key)
        fmt.Println(list.Node.Value)
    }


/*
    upstreamUrl := "http://127.0.0.1:9090"
    remote, err := url.Parse(upstreamUrl)
    if err != nil {
        panic(err)
    }


    http.Handle("/", r)
    http.ListenAndServe(":8000", r)
*/
}
/*
func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
    return func(responseWriter http.ResponseWriter, request *http.Request) {
        request.URL.Path = mux.Vars(request)["rest"]
        t := time.Now()
        p.ServeHTTP(responseWriter, request)
        fmt.Println(time.Since(t).Seconds())
    }
}
*/
