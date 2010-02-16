package main

import (
        "fmt";
        "strings";
        "flag";
)

func loadlist(m map[int]string) {
    c := make(chan string);
    tf := make(chan int);
    DBI:=newDBI()
    ts := 3;
    jobsplit:=len(m)/ts;
    jobmod:=len(m)%ts;
    for i := 0; i < ts ; i++ {
        if jobmod>0 {
            go DBI.urlparser(i,jobsplit+1,c,tf);
            jobmod--;
        }else{
            go DBI.urlparser(i,jobsplit,c,tf);
        }
    }
    // Sending jobs to each go channel.
    for _, url := range m {
        c <- url;
    }
    // Waiting for all of those threads are finsh.
    for i := 0; i < ts ; i++ {
        <-tf;
    }
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

