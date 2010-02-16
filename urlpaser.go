package main

import (
        "regexp";
        "strings";
        "strconv";
        "fmt";
        "time";
        "mysql";
        "./configfile";
)

func printamule(el string,p *configfile.ConfigFile){
    ars,_:=p.GetString("amule","ARS")
    arp,_:=p.GetString("amule","ARP")
    arps,_:=p.GetString("amule","ARPS")
    fmt.Printf("amulecmd --host=%s -p %s -P %s -c \"add %s\"\n",ars,arp,arps,el)
}

func urlparser(id int,size int,c chan string,tf chan int) {
    ed2k,_ := regexp.Compile("href=\"ed2k://");
    re,_ := regexp.Compile("<([^>]|\n)*>|\t|\r");
    p, err := configfile.ReadConfigFile("config.cfg");
    if err != nil { fmt.Printf("Load error: %s\n",err); }
    if err != nil { fmt.Printf("Connection error: %s\n",err); }
    db_name,_ := p.GetString("default","DBN")
    DBIP,_:=p.GetString("default","DBIP")
    DBA,_:=p.GetString("default","DBA")
    DBP,_:=p.GetString("default","DBP")
    dbh, err := mysql.Connect("tcp", "", DBIP,DBA,DBP, db_name)
    if err != nil { fmt.Printf("DB error: %s\n",err); }
    coulmlist := "(`scheme` ,`type` ,`filename` ,`filesize` ,`hash`,`ori`,`rctime`) VALUES "
    sqlc:=fmt.Sprintf("%s%s%s%s","INSERT INTO `",db_name,"`.`ed2k` ",coulmlist);

    //coulmlist := "(`scheme` ,`type` ,`filename` ,`filesize` ,`hash`,`ori`,`rctime`) VALUES (?,?,?,?,?,?,FROM_UNIXTIME(?))"
    //sqlcommand:=fmt.Sprintf("%s%s%s%s","INSERT INTO `",db_name,"`.`ed2k` ",coulmlist);
    //stmt, err := conn.Prepare(sqlcommand);
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
                insevalue:=fmt.Sprintf("%s ('%s','%s','%s','%s','%s','%s',FROM_UNIXTIME(%s))",
                sqlc,stsli[0],stsli[1],stsli[2],stsli[3],stsli[4],edurl,strconv.Itoa64(time.Seconds()))
                _, e := dbh.Query(insevalue)
                //_, e := conn.Execute(stmt, stsli[0],stsli[1],stsli[2],stsli[3],stsli[4],edurl,strconv.Itoa64(time.Seconds()));
                if e == nil {
                    if p.HasSection("amule"){
                        printamule(edurl,p)
                    }else{
                        fmt.Printf("%s",edurl)
                    }
                }
            }
        }
        fmt.Printf("#ID: %d finsh job %d\n",id,i);
    }
    dbh.Quit();
    tf<-1;
}

