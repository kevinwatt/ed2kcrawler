package main

import (
        "mysql";
        "regexp";
        "strings";
        "strconv";
        "fmt";
        "time";
        "io/ioutil";
)

func loadconfig() (config map[string]string){
    b, err :=ioutil.ReadFile("config.inc");
    config = make(map[string]string);
    if err==nil {
        space,_ := regexp.Compile(" ");
        cont:=string(b);
        cont=space.ReplaceAllString(cont,"");
        configlist:=strings.Split(cont, "\n", 0);
        for i := 0; i < len(configlist); i++ {
            if !strings.HasPrefix(configlist[i],";") || !strings.HasPrefix(configlist[i],"#") {
                kv:=strings.Split(configlist[i],"=",0);
                if len(kv)>1{
                    config[kv[0]]=kv[1];
                }
            }
        }
    }
    return;
}

func urlparser(id int,size int,c chan string,tf chan int) {
    ed2k,_ := regexp.Compile("href=\"ed2k://");
    re,_ := regexp.Compile("<([^>]|\n)*>|\t|\r");
    kv:=loadconfig();
    conn, err := mysql.Open(kv["DB"]);
    if err != nil { fmt.Printf("Connection error: %s\n",err); }
    stmt, err := conn.Prepare("INSERT INTO `godba`.`ed2k` (`scheme` ,`type` ,`filename` ,`filesize` ,`hash`,`ori`,`rctime`) VALUES (?,?,?,?,?,?,FROM_UNIXTIME(?))");
    if err != nil { fmt.Printf("Command error: %s\n",err); }
    for i := 0; i < size; i++ {
        url := <-c;
        pas := Get(url);
        pas = ed2k.ReplaceAllString(pas,">\ned2k://");
        pas = re.ReplaceAllString(pas,"");
        pasarray := strings.Split(pas,"\n",0);
        for i := 0; i < len(pasarray); i++ {
            if strings.HasPrefix(pasarray[i],"ed2k://") {
                edurl := pasarray[i][0:strings.Index(pasarray[i],"\"")];
                edurl = URLUnescape(edurl);
                stsli := strings.Split(edurl,"|",0)
                _, e := conn.Execute(stmt, stsli[0],stsli[1],stsli[2],stsli[3],stsli[4],edurl,strconv.Itoa64(time.Seconds()));
                if e == nil {
                    fmt.Printf("%s\n",edurl);
                }
            }
        }
        fmt.Printf("ID: %d finsh job %d\n",id,i);
    }
    stmt.Close(); conn.Close();
    tf<-1;
}

