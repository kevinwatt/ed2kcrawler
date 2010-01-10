package main
import (
        "http"
        "strings"
        "fmt"
        "io/ioutil"
)

func Get(url string) string {
    n := new(http.Response);
    n,_,err := http.Get(http.URLEscape(url));
    var b []byte;
    if err!=nil {
        fmt.Println("error to read url ",err,url);
    }else{
        b , _ = ioutil.ReadAll(n.Body)
        n.Body.Close();
    }
    return string(b);
}

func URLUnescape(edurl string) string {
    edurl,_ = http.URLUnescape(edurl);
    return edurl;
}


