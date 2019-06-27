package httputil

import (
	"compress/gzip"
	"github.com/dsnet/compress/brotli"
	"io"
	"io/ioutil"
	"net/http"
)

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
