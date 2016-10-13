package store

import (
    "io"
    "time"
    "strings"
    "encoding/json"
    "github.com/microcats/hellgate/backend"
)

type ApiInfo struct {
    Prefix, UpstreamUrl string
    CreateAt time.Time
}

func GetApiInfo() (c *backend.Client, map[int]*ApiInfo, error) {
    list, _ := c.Get("/hellgate/apis")
    //{"prefix":"test1","upstreamUrl":"http://127.0.0.1:9090", "createAt":"2016-02-01 15:11:22"}
    apiInfos := make(map[int]*ApiInfo, 0)
    for key, value := range list.Node.Nodes {
        apiInfo := new(ApiInfo)
        info, _ := c.Get(value.Key)
        decode := json.NewDecoder(strings.NewReader(info.Node.Value))
        if err := decode.Decode(&apiInfo); err == io.EOF {
            return nil, err
        } else if err != nil {
            return nil, err
        }

        apiInfos[key] = apiInfo
    }

    return apiInfos, nil
}
