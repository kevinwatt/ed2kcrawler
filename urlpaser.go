package main

import (
        "regexp";
        "strings";
	"url";
	"fmt";
)

type URLParse struct {
    id              int
    size            int
}

func (up *URLParse) urlparser(c chan string, tf chan map[string]string) {
    ed2k,_ := regexp.Compile("href=\"ed2k://");
    re,_ := regexp.Compile("<([^>]|\n)*>|\t|\r");
    parsedlink:= make(map[string]string)
    for i := 0; i < up.size; i++ {
        if pas,err := Get(<-c); err == nil {
            pas = ed2k.ReplaceAllString(pas,">\ned2k://");
            pas = re.ReplaceAllString(pas,"\n");
            pasarray := strings.Split(pas,"\n");
            for is := 1; is < len(pasarray); is++ {
                if strings.HasPrefix(pasarray[is],"ed2k://") {
                    stringindex:=strings.Index(pasarray[is],"\"")
                    var edurl string
                    if stringindex < 1{
                        edurl = pasarray[is];
                        edurl,_ = url.QueryUnescape(edurl);
                    }else{
                        edurl = pasarray[is][0:stringindex];
                        edurl,_ = url.QueryUnescape(edurl);
                    }
                    spedurl := strings.Split(edurl,"|")
                    if len(spedurl) > 5 && len(spedurl[4]) > 20  {
                        key:=spedurl[4]
                        parsedlink[key]=edurl;
                    }
                }
            }
        }else{
	    fmt.Printf("can't open url; err=%s\n",  err.String())
	}
    }
    tf<-parsedlink;
    return;
}

