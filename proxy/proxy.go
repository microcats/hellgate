package proxy

import (
    "log"
    //"fmt"
    "time"
    //"net/url"
    //"net/http"
    //"net/http/httputil"
    //"github.com/gorilla/mux"
    "context"
    "github.com/coreos/etcd/client"
)


func NewMultipleHostReverseProxy() {
    cfg := client.Config{
        Endpoints:               []string{"http://127.0.0.1:2379"},
        Transport:               client.DefaultTransport,
        // set timeout per request to fail fast when the target endpoint is unavailable
        HeaderTimeoutPerRequest: time.Second,
    }

    c, err := client.New(cfg)
    if err != nil {
        log.Fatal(err)
    }
    kapi := client.NewKeysAPI(c)

    resp, err := kapi.Get(context.Background(), "/hellgate/apis", &client.GetOptions{Recursive: true})
    if err != nil {
        log.Fatal(err)
    } else {
        log.Println(resp.Node.Nodes)

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
