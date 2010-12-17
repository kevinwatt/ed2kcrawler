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
    ed2kurldb       *URLStore
    size            int
    id              int
    tf              chan int
}

func printamule(el string,p *configfile.ConfigFile){
    ars,_:=p.GetString("amule","ARS")
    arp,_:=p.GetString("amule","ARP")
    arps,_:=p.GetString("amule","ARPS")
    fmt.Printf("amulecmd --host=%s -p %s -P %s -c \"add %s\"\n",ars,arp,arps,el)
}

func (up *URLParse) urlparser(c chan string) {
    ed2k,_ := regexp.Compile("href=\"ed2k://");
    re,_ := regexp.Compile("<([^>]|\n)*>|\t|\r");
    //ed2kurldb:=NewURLStore("ed2kurl.gmap")
    for i := 0; i < up.size; i++ {
        url := <-c;
        pas := Get(url);
        pas = ed2k.ReplaceAllString(pas,">\ned2k://");
        pas = re.ReplaceAllString(pas,"\n");
        pasarray := strings.Split(pas,"\n",-1);
        //fmt.Printf("#ID: %d take the job %s size: %d\n",id,url,len(pasarray));
        for is := 1; is < len(pasarray); is++ {
            if strings.HasPrefix(pasarray[is],"ed2k://") {
                /*edurl := pasarray[i][0:strings.Index(pasarray[is],"\"")];
                edurl = URLUnescape(edurl);
                fmt.Printf("%s\n",edurl)
                */
                stringindex:=strings.Index(pasarray[is],"\"")
                var edurl string;
                if stringindex < 1{
                    edurl = pasarray[is];
                    edurl = URLUnescape(edurl);
                    //key = strings.Split(pasarray[is],"|",-1)[4]
                }else{
                    edurl = pasarray[is][0:stringindex];
                    edurl = URLUnescape(edurl);
                    //key = strings.Split(edurl,"|",-1)[4]
                }
                spedurl := strings.Split(edurl,"|",-1)
                if len(spedurl) > 5 && len(spedurl[4]) > 20  {
                    key:=spedurl[4]
                    var getedurl string;
                    if up.ed2kurldb.Get(&key,&getedurl); len(getedurl) < 1 {
                        fmt.Printf("getedurl < 1 ::: %s %s\n",edurl,key);
                        //ed2kurldb.Get(&key,&getedurl);
                        //fmt.Printf("%s %s\n",getedurl,key);
                        up.ed2kurldb.Put(&edurl,&key);
                    }
                }
            }
        }
    }
    fmt.Printf("%d finsh\n",up.id);
    up.tf<-1;
    return;
}

