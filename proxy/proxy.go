package proxy

import (
    "fmt"
    "time"
    "net/url"
    "net/http"
    "net/http/httputil"
    "github.com/gorilla/mux"
    "github.com/microcats/hellgate/backend"
    "github.com/microcats/hellgate/store"
)


func NewMultipleHostReverseProxy(c *backend.Client) (*mux.Router, error) {
    apis, err := store.GetApiInfo(c)
    if err != nil {
        panic(err)
    }

    r := mux.NewRouter()
    for _, api := range apis {
        remote, err := url.Parse(api.UpstreamUrl)
        if err != nil {
            return nil, err
        }

        proxy := httputil.NewSingleHostReverseProxy(remote)
        path := fmt.Sprintf("/%s/{rest:.*}", api.Prefix)
        r.HandleFunc(path, handler(proxy))
    }

    return r, nil
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
    return func(responseWriter http.ResponseWriter, request *http.Request) {
        request.URL.Path = mux.Vars(request)["rest"]
        t := time.Now()
        p.ServeHTTP(responseWriter, request)
        fmt.Println(time.Since(t).Seconds())
    }
}
