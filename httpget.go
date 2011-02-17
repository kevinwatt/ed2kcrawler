package main

import (
        "http"
        "strings"
        "io"
        "io/ioutil"
        "os"
)

type readClose struct {
    io.Reader
    io.Closer
}

func hasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func Get(url string) (string, os.Error) {
    //var b []byte;
    resp,_,err := http.Get(url)
    b , _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close();
    return string(b),err
}

func URLUnescape(edurl string) string {
    edurl , _ = http.URLUnescape(edurl);
    return edurl;
}


