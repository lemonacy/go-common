package hoo

import (
	"compress/gzip"
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

func ReadResponse(resp *http.Response) string {
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
	text := string(body)
	return text
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
