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
    /*
    s := &URLParse{
        ed2kurldb: NewURLStore("ed2kurl.gmap"),
    }
    */
    ts := 3
    jobsplit:=len(m)/ts;
    jobmod:=len(m)%ts;
    for i := 0; i < ts ; i++ {
        if jobmod>0 {
            s := &URLParse{ ed2kurldb: NewURLStore("ed2kurl.gmap") }
            s.id=i;
            s.size=jobsplit+1
            go s.urlparser(c,tf);
            jobmod--;
        }else{
            s := &URLParse{ ed2kurldb: NewURLStore("ed2kurl.gmap") }
            s.id=i;
            s.size=jobsplit
            go s.urlparser(c,tf);
        }
    }

    for _, url := range m {
        c <- url;
    }
    for i := 0; i < ts ; i++ {
        print(<-tf);
    }
}

func help(){
    fmt.Printf("%s\n","ed2kcrawler v0.0.2pre");
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

