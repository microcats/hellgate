package store

import (
    "io"
    "time"
    "strings"
    "encoding/json"
    "github.com/microcats/hellgate/backend"
)

type ReverseProxy struct {
    Prefix, UpstreamUrl string
    CreateAt time.Time
}

func GetReverseProxy() (c *backend.Client, map[int]*ReverseProxy, error) {
    result, _ := c.Get("/hellgate/apis")
    //{"prefix":"test1","upstreamUrl":"http://127.0.0.1:9090", "createAt":"2016-02-01 15:11:22"}
    reverseProxys := make(map[int]*ReverseProxy, 0)
    for key, value := range result.Node.Nodes {
        reverseProxy := new(ReverseProxy)
        list, _ := c.Get(value.Key)
        decode := json.NewDecoder(strings.NewReader(list.Node.Value))
        if err := decode.Decode(&reverseProxy); err == io.EOF {
            return nil, err
        } else if err != nil {
            return nil, err
        }

        reverseProxys[key] = reverseProxy
    }

    return reverseProxys, nil
}
