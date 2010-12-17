package main

import (
    "os";
    "strings";
    "io/ioutil";
)

func loadvv(fname string) (ul map[int]string,err os.Error) {
    b, err :=ioutil.ReadFile(fname);
    ul = make(map[int]string);
    if err==nil {
        cont:=string(b);
        slist:=strings.Split(cont, "\n", -1);
        for p,i := 0,0; i < len(slist); i++ {
          if strings.HasPrefix(slist[i],"http://") {
              ul[p]=slist[i];
              p++;
          }
        }
    }
    return;
}


