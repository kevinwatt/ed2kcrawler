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

func ContGet(url string) (string, os.Error) {

    var err os.Error;
    resp, err := http.Get(url);
    if err == nil {
            b , err := ioutil.ReadAll(resp.Body)
            resp.Body.Close();
            return string(b),err
    }
    return "error",err
}

