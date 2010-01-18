package main

import (
        "http"
        "strings"
        "fmt"
        "io/ioutil"
        "net"
        "strconv"
        "os"
        "io"
        "bufio"
        "encoding/base64"
)

type readClose struct {
    io.Reader
    io.Closer
}

func hasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func send(req *http.Request) (resp *http.Response, err os.Error) {
    addr := req.URL.Host
    if !hasPort(addr) {
        addr += ":http"
    }
    info := req.URL.Userinfo
    if len(info) > 0 {
        enc := base64.URLEncoding
        encoded := make([]byte, enc.EncodedLen(len(info)))
        enc.Encode(encoded, strings.Bytes(info))
        if req.Header == nil {
            req.Header = make(map[string]string)
        }
        req.Header["Authorization"] = "Basic " + string(encoded)
    }
    conn, err := net.Dial("tcp", "", addr)
    if err != nil {
        return nil, err
    }

    err = req.Write(conn)
    if err != nil {
        conn.Close()
        return nil, err
    }

    reader := bufio.NewReader(conn)
    resp, err = http.ReadResponse(reader)
    if err != nil {
        conn.Close()
        return nil, err
    }

    r := io.Reader(reader)
    if v := resp.GetHeader("Content-Length"); v != "" {
        n, _ := strconv.Atoi64(v)
        r = io.LimitReader(r, n)
    }
    resp.Body = readClose{r, conn}
    return
}

func Get(url string) string {
    n := new(http.Response);
    //n,_,err := http.Get(http.URLEscape(url));
    var req http.Request
    req.URL, _ = http.ParseURL(http.URLEscape(url));
    req.UserAgent = "Mozilla/5.0 ed2kcrawler"
    n,err := send(&req);

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
    edurl , _ = http.URLUnescape(edurl);
    return edurl;
}


