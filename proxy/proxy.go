package proxy

import (
    "fmt"
    "time"
    "net/url"
    "net/http"
    "net/http/httputil"
    //"encoding/json"
    "io/ioutil"
)

type upstream struct {
    prefix  string
    host string
}

func New(prefix string, host string) *upstream {
    upstream := new(upstream)
    upstream.prefix = prefix
    upstream.host = host
    return upstream
}

func (u *upstream) Register() {
    http.HandleFunc(u.getPrefix(), u.handler)
}

func (u *upstream) getPrefix() string {
    return fmt.Sprintf("/%s/", u.prefix)
}

func (u *upstream) handler(w http.ResponseWriter, request *http.Request) {
    upstream, err := url.Parse(u.host)
    if err != nil {
        panic(err)
    }


    if (request.Header.Get("Content-Type") == "application/x-www-form-urlencoded") {
        request.ParseForm()
    } else {
        request.ParseMultipartForm(1024)
    }
    

    body, _ := ioutil.ReadAll(request.Body)
    fmt.Printf("%s", body)
    fmt.Println(request.Form)
    fmt.Println(request.PostForm)
    fmt.Println(request.MultipartForm)


    //request.ParseForm()

    //r.ParseForm()
    //r.ParseMultipartForm(1024)
    //application/json
    //multipart/form-data
    //application/x-www-form-urlencoded
    //fmt.Println(r.Header.Get("Content-Type"))

    for k, v := range request.Header {
        fmt.Printf("key[%s] value[%s]\n", k, v)
    }

    fmt.Println(request.Method)
    fmt.Println(request.Host)
    fmt.Println(request.TransferEncoding)
    fmt.Println(request.ContentLength)
    fmt.Println(request.URL)
    fmt.Println(request.Proto)
    fmt.Println(request.ProtoMajor)
    fmt.Println(request.ProtoMinor)


    proxy := http.StripPrefix(u.getPrefix(), httputil.NewSingleHostReverseProxy(upstream))
    start := time.Now()
    proxy.ServeHTTP(w, request)
    responseTime := time.Since(start).Seconds()
    fmt.Println(request.RemoteAddr)
    fmt.Println(request.RequestURI)
    fmt.Println(responseTime)

}
