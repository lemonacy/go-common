package hoo

import (
    "compress/gzip"
    "fmt"
    "github.com/dsnet/compress/brotli"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
)

func NewCookieClient() *http.Client {
    jar, err := cookiejar.New(nil)
    if err != nil {
        panic(err)
    }

    client := http.Client{Jar: jar}
    return &client
}

func NewDownloadClient() *http.Client {
    // Customize the Transport to have larger connection pool
    defaultRoundTripper := http.DefaultTransport
    defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
    if !ok {
        panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
    }
    defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
    defaultTransport.MaxIdleConns = 1
    defaultTransport.MaxIdleConnsPerHost = 1

    return &http.Client{
        Transport: &defaultTransport,
        Timeout:   60 * 1000 * 1000 * 1000,
    }
}

func ReadResponseString(resp *http.Response) string {
    text := string(ReadResponse(resp))
    return text
}

func ReadResponse(resp *http.Response) []byte {
    var reader io.ReadCloser
    var err error
    switch resp.Header.Get("Content-Encoding") {
    case "gzip":
        reader, err = gzip.NewReader(resp.Body)
        if err != nil {
            panic(err)
        }
        defer reader.Close()
    case "br":
        reader, err = brotli.NewReader(resp.Body, nil)
        defer reader.Close()
    default:
        reader = resp.Body
    }
    body, err := ioutil.ReadAll(reader)
    if err != nil {
        panic(err)
    }
    return body
}

func PrintReqHeader(req *http.Request) {
    printHeader(req.Header)
}

func PrintRespHeader(resp *http.Response) {
    printHeader(resp.Header)
}

func printHeader(header http.Header) {
    for name, values := range header {
        for _, value := range values {
            log.Println(name + ": " + value)
        }
    }
}
