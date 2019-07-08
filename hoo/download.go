/**
 * Copyright (C) 1995-2019 Seasun Entertainment 西山居 版权所有;
 *
 * Author: dengkuadong <dengkuadong@kingsoft.com>
 * Since: 2019-07-04
 */

package hoo

import (
    "compress/gzip"
    "errors"
    "github.com/dsnet/compress/brotli"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

type Downloader interface {
    Download(url, storedir string) (file string, err error)
    DownloadBytes(url string) ([]byte, error)
}

type downloader struct {
    client *http.Client
}

func (c *downloader) Download(url, storedir string) (file string, err error) {
    log.Printf("start to downlaod %s", url)
    filename := filepath.Base(url)
    path := filepath.Join(storedir, filename)
    b, err := FileExists(path)
    PanicOnErr(err)
    if b {
        log.Printf("%s already downloaded, skip.\n", filepath.Base(path))
        return path, nil
    }

    success := false
    for i := 0; i < 5; i++ {
        resp, err := c.client.Get(url)
        if err != nil {
            continue
        }
        var reader io.ReadCloser
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
            continue
        }

        saveto, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 644)
        PanicOnErr(err)
        saveto.Write(body)
        saveto.Close()

        success = true
        break
    }

    if success {
        log.Println("downloaded")
        return path, nil
    } else {
        return "", errors.New("Failed to download " + url)
    }
}

func (c *downloader) DownloadBytes(url string) ([]byte, error) {
    log.Printf("start to downlaod %s", url)

    success := false

    var body []byte
    for i := 0; i < 5; i++ {
        //req, err := http.NewRequest("GET", url, nil)
        //PanicOnErr(err)
        //req.Header.Set("Host", Host(url))
        //req.Host = Host(url)
        //resp, err := c.client.Do(req)
        resp, err := c.client.Get(url)
        if err != nil {
            log.Println(err.Error())
            continue
        }
        log.Println(resp.Status)
        var reader io.ReadCloser
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
        body, err = ioutil.ReadAll(reader)
        if err != nil {
            continue
        }

        success = true
        break
    }

    if success || len(body) == 0 {
        log.Println("downloaded")
        return body, nil
    } else {
        return nil, errors.New("Failed to download " + url)
    }
}

func NewDownloader() Downloader {
    d := downloader{}
    d.client = NewDownloadClient()
    return &d
}
