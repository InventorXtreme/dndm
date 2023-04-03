package main

import "fmt"
import "github.com/chzyer/readline"
import "regexp"
import "strings"

type mapobj struct {
    val string
    sub map[string]*mapobj
}

func SToMap(s string,m mapobj) *mapobj {
    r := regexp.MustCompile("^[a-zA-Z0-9]*/")
    a := r.FindString(s)
    path := strings.TrimSuffix(a, "/")
    newa := strings.TrimPrefix(s,a)
    //fmt.Println(path)
    if(len(newa) > 0){
        return SToMap(newa, *m.sub[path])
    } else {
        return m.sub[path]
    }
}

func MakeMapObj(x string) *mapobj {
    ipm := new(mapobj)
    ipm.val = x
    ipm.sub = make(map[string]*mapobj)
    return ipm
}





func main(){
    maintime := MakeMapObj("sub")
    maintime.sub["players"] = MakeMapObj("sub")

    pc := readline.NewPrefixCompleter(readline.PcItem("hehe"),readline.PcItem("hoho",readline.PcItem("haha")),readline.PcItem("hehehe"))

    l, _ := readline.NewEx(&readline.Config{Prompt: "aa ", AutoComplete: pc})
    
    readline.SetAutoComplete(pc)
    b, _ := l.Readline()
    
    fmt.Println(b)

    main := MakeMapObj("sub")
    main.sub["hehe"] = MakeMapObj("sub")
    main.sub["hehe"].sub["test"] = MakeMapObj("data")
    //fmt.Println(main.sub["hehe"].val)
    fmt.Println(SToMap("hehe/test/",*main).val)
}
