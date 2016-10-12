package backend

import (
    "time"
    "context"
    "net"
    "net/http"
    "io/ioutil"
    "crypto/tls"
    "crypto/x509"
    "github.com/coreos/etcd/client"
)

type Client struct {
    client client.KeysAPI
}

func NewEtcdClient(machines []string, cert, key, caCert string, basicAuth bool, username string, password string) (*Client, error) {
    var c client.Client
    var kapi client.KeysAPI
    var err error
    var transport = &http.Transport{
        Proxy: http.ProxyFromEnvironment,
        Dial: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
        }).Dial,
        TLSHandshakeTimeout: 10 * time.Second,
    }

    tlsConfig := &tls.Config{
        InsecureSkipVerify: false,
    }

    cfg := client.Config{
        Endpoints:               machines,
        HeaderTimeoutPerRequest: time.Duration(3) * time.Second,
    }

    if basicAuth == true {
        cfg.Username = username
        cfg.Password = password
    }

    if caCert != "" {
        certBytes, err := ioutil.ReadFile(caCert)
        if err != nil {
            return &Client{kapi}, err
        }

        caCertPool := x509.NewCertPool()
        ok := caCertPool.AppendCertsFromPEM(certBytes)

        if ok {
            tlsConfig.RootCAs = caCertPool
        }
    }

    if cert != "" && key != "" {
        tlsCert, err := tls.LoadX509KeyPair(cert, key)
        if err != nil {
            return &Client{kapi}, err
        }
        tlsConfig.Certificates = []tls.Certificate{tlsCert}
    }

    transport.TLSClientConfig = tlsConfig
    cfg.Transport = transport

    c, err = client.New(cfg)
    if err != nil {
        return &Client{kapi}, err
    }

    kapi = client.NewKeysAPI(c)
    return &Client{kapi}, nil
}

func (c *Client) Get(key string) (*client.Response, error) {
    return c.client.Get(context.Background(), key, &client.GetOptions{
        Recursive: true,
        Sort:      true,
        Quorum:    true,
    })
}
