package store

import (
    "io"
    //"fmt"
    "time"
    "strings"
    "encoding/json"
    "github.com/microcats/hellgate/backend"
)

type ApiInfo struct {
    Prefix, UpstreamUrl string
    CreateAt time.Time
}

func getApiInfo(c *backend.Client, key string, channel chan *ApiInfo, e chan error) {
    apiInfo := new(ApiInfo)
    info, _ := c.Get(key)
    decode := json.NewDecoder(strings.NewReader(info.Node.Value))
    if err := decode.Decode(&apiInfo); err == io.EOF {
        e <- err
    } else if err != nil {
        e <- err
    }

    channel <- apiInfo
    close(channel)
    close(e)
}

func GetApiInfo(c *backend.Client) (map[int]*ApiInfo, error) {
    list, _ := c.Get("/hellgate/apis")
    //{"prefix":"test1","upstreamUrl":"http://127.0.0.1:9090", "createAt":"2016-06-09T07:12:17Z"}
    apiInfos := make(map[int]*ApiInfo, 0)
    for key, value := range list.Node.Nodes {
        channel := make(chan *ApiInfo, 1)
        err := make(chan error, 1)
        go getApiInfo(c, value.Key, channel, err)
        select {
        case e := <-err:
            if e != nil {
                return nil, e
            }
        }

        apiInfos[key] = <- channel
    }

    return apiInfos, nil
}
