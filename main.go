package main

import (
        "fmt";
        "strings";
        "flag";
)

func loadlist(m map[int]string) {
    //store:=NewURLStore("store.gmap")
    c := make(chan string);
    tf := make(chan int);
    s := &URLParse{
        ed2kurldb: NewURLStore("ed2kurl.gmap"),
    }

    /*
    jobsplit:=len(m)/ts;
    jobmod:=len(m)%ts;
    for i := 0; i < ts ; i++ {
        if jobmod>0 {
            s.id=i;
            s.size=jobsplit+1
            s.tf=tf
            go s.urlparser(c);
            jobmod--;
        }else{
            s.id=i;
            s.size=jobsplit
            s.tf=tf
            go s.urlparser(c);
        }
    }
    */
    s.id=0;
    s.size=len(m)
    s.tf=tf
    go s.urlparser(c);

    for _, url := range m {
        c <- url;
        //var h hash.Hash = md5.New()
        //h.Write([]byte(url))
        //urlmd5:=h.Sum()
        //urlmd5:=fmt.Sprintf("%x", h.Sum())
        //store.Put(&url, &urlmd5);
    }
    print(<-s.tf);
}

func help(){
    fmt.Printf("%s\n","ed2kcrawler v0.0.1pre");
    fmt.Printf("%s\n","Usage: ed2kcrawler [Option]... [URL]...");
    fmt.Printf("%s\n","Commands:");
    flag.PrintDefaults();
}

func main() {
    var Loadpagelist = flag.Bool("l", false, "\tLoad the ed2klink page url list")
    flag.Parse();
    if *Loadpagelist {
        listfilename:=flag.Arg(0);
        m,err:=loadvv(listfilename);
        if err==nil {
            loadlist(m);
        } else {
            fmt.Printf("File %s not exist.\n",listfilename);
        }
    }else{
        if flag.NArg() >0 {
            m := make(map[int]string);
            for p,i := 0,0; i < flag.NArg(); i++ {
                if strings.HasPrefix(flag.Arg(i),"http://") {
                    m[p]=flag.Arg(i);
                    p++;
                }
            }
            loadlist(m);
        }else{
            help();
        }
    }
}

