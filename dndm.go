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


func SetValToPath(s string, val string, m mapobj) *mapobj {
	r := regexp.MustCompile("^[a-zA-Z0-9]*/")
    a := r.FindString(s)
    path := strings.TrimSuffix(a, "/")
    newa := strings.TrimPrefix(s,a)
    //fmt.Println(path)
    if(len(newa) > 0){
        _, ok := m.sub[path]
		if (!ok) {
			m.sub[path] = MakeMapObj("sub")
		}
        return SetValToPath(newa, val, *m.sub[path])
    } else {
		m.sub[path] = MakeMapObj(val)
		//m.val = val
        return &m
    } 
}

func PathFormat(toformat string) string {
	if(toformat[len(toformat)-1:] != "/") {
		return ( toformat + "/")
	} else {
		return toformat
	}
}

func main(){
    maintime := MakeMapObj("sub")
    maintime.sub["players"] = MakeMapObj("sub")
    SetValToPath("home/","test",*maintime)
	fmt.Println(SToMap("home/", *maintime).val)
    pc := readline.NewPrefixCompleter(readline.PcItem("set"),readline.PcItem("get"),readline.PcItem("quit"))

    l, _ := readline.NewEx(&readline.Config{Prompt: "? ", AutoComplete: pc})
   
	for true {
		command, _ := l.Readline()
		if command == "quit" {
			break
		}
		

	} 
    readline.SetAutoComplete(pc)
    //b, _ := l.Readline()
    
    //fmt.Println(b)

    main := MakeMapObj("sub")
    main.sub["hehe"] = MakeMapObj("sub")
    main.sub["hehe"].sub["test"] = MakeMapObj("data")
    //fmt.Println(main.sub["hehe"].val)
    fmt.Println(SToMap("hehe/",*main).val)
}
