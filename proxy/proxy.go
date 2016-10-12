package proxy

import (
    //"log"
    "io"
    "fmt"
    "time"
    "strings"
    "encoding/json"
    "net/url"
    "net/http"
    "net/http/httputil"
    "github.com/gorilla/mux"
    "github.com/microcats/hellgate/backend"
)


type apiList struct {
    Prefix, UpstreamUrl string
    CreateAt time.Time
}

func getApiList() (map[int]*apiList, error) {
    machines := []string{"http://127.0.0.1:2379"}
    etcd, _ := backend.NewEtcdClient(machines, "", "", "", false, "", "")
    result, _ := etcd.Get("/hellgate/apis")
    //{"prefix":"test1","upstreamUrl":"http://127.0.0.1:9090", "createAt":"2016-02-01 15:11:22"}
    apiLists := make(map[int]*apiList, 0)
    for key, value := range result.Node.Nodes {
        apiList := new(apiList)
        list, _ := etcd.Get(value.Key)
        decode := json.NewDecoder(strings.NewReader(list.Node.Value))
        if err := decode.Decode(&apiList); err == io.EOF {
            return nil, err
        } else if err != nil {
            return nil, err
        }

        apiLists[key] = apiList
    }

    return apiLists, nil
}

func NewMultipleHostReverseProxy() (*mux.Router, error) {
    apis, err := getApiList()
    if err != nil {
        panic(err)
    }

    r := mux.NewRouter()
    for _, api := range apis {
        fmt.Println(api.UpstreamUrl)
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
