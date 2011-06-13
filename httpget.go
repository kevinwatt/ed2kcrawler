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
    var err os.Error;
    if resp, err := http.Get(url); err == nil {
	    b , err := ioutil.ReadAll(resp.Body)
	    resp.Body.Close();
	    return string(b),err
    }
    return "error",err
}

func URLUnescape(edurl string) string {
    edurl , _ = http.URLUnescape(edurl);
    return edurl;
}


