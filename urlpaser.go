package main

import (
        "regexp";
        "strings";
        //"strconv";
        "fmt";
        //"time";
        "./configfile";
)

type URLParse struct {
    id              int
    size            int
}

func printamule(el string,p *configfile.ConfigFile){
    ars,_:=p.GetString("amule","ARS")
    arp,_:=p.GetString("amule","ARP")
    arps,_:=p.GetString("amule","ARPS")
    fmt.Printf("amulecmd --host=%s -p %s -P %s -c \"add %s\"\n",ars,arp,arps,el)
}

func (up *URLParse) urlparser(c chan string, tf chan map[string]string) {
    ed2k,_ := regexp.Compile("href=\"ed2k://");
    re,_ := regexp.Compile("<([^>]|\n)*>|\t|\r");
    parsedlink:= make(map[string]string)
    for i := 0; i < up.size; i++ {
        pas := Get(<-c);
        pas = ed2k.ReplaceAllString(pas,">\ned2k://");
        pas = re.ReplaceAllString(pas,"\n");
        pasarray := strings.Split(pas,"\n",-1);
        for is := 1; is < len(pasarray); is++ {
            if strings.HasPrefix(pasarray[is],"ed2k://") {
                stringindex:=strings.Index(pasarray[is],"\"")
                var edurl string
                if stringindex < 1{
                    edurl = pasarray[is];
                    edurl = URLUnescape(edurl);
                }else{
                    edurl = pasarray[is][0:stringindex];
                    edurl = URLUnescape(edurl);
                }
                spedurl := strings.Split(edurl,"|",-1)
                if len(spedurl) > 5 && len(spedurl[4]) > 20  {
                    key:=spedurl[4]
                    parsedlink[key]=edurl;
                }
            }
        }
    }
    tf<-parsedlink;
    return;
}

