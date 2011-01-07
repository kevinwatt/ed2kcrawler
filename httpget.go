package main

import (
        "http"
        "strings"
        "io"
        "io/ioutil"
)

type readClose struct {
    io.Reader
    io.Closer
}

func hasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func Get(url string) string {
    var b []byte;
    resp,_,_ := http.Get(url)
    b , _ = ioutil.ReadAll(resp.Body)
    resp.Body.Close();

    return string(b);
}

func URLUnescape(edurl string) string {
    edurl , _ = http.URLUnescape(edurl);
    return edurl;
}


