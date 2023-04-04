package main

import "fmt"
import "github.com/chzyer/readline"
import "regexp"
import "strings"

type mapobj struct {
    val string
    sub map[string]*mapobj
}

type completeholder struct {
	m *mapobj 
}

func SToMap(s string,m mapobj) *mapobj {
    r := regexp.MustCompile("^[a-zA-Z0-9]*/")
    a := r.FindString(s)
    path := strings.TrimSuffix(a, "/")
    newa := strings.TrimPrefix(s,a)
    //fmt.Println(path)
    if(len(newa) > 0){
        _,ok := m.sub[path]
        if(!ok) {
            return MakeMapObj("Item not found")
        }
        return SToMap(newa, *m.sub[path])
    } else {
	_,ok := m.sub[path]
        if(!ok){
            return MakeMapObj("Item not found")
        }
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
	e := strings.TrimSuffix(toformat, " ")
	if(e[len(e)-1:] != "/") {
		return ( e + "/")
	} else {
		return e
	}
}

func ListAllPaths(prefix string,m *mapobj) []string {
	var l []string
	for k, v := range m.sub {
		l = append(l,(prefix + k + "/"))
		if len(v.sub) > 0 {
			e := ListAllPaths((prefix + k + "/"),v)
			for _,v := range e {
				l = append(l,v)
			}
		}
	}
	return l
}

func (e completeholder ) compthing(a string) []string {
	return ListAllPaths("",e.m)
}

func main(){
    maintime := MakeMapObj("sub")
    maintime.sub["players"] = MakeMapObj("sub")
    SetValToPath("home/","test",*maintime)
	fmt.Println(SToMap("home/", *maintime).val)
    a := completeholder{maintime}
    pc := readline.NewPrefixCompleter(readline.PcItem("set",readline.PcItemDynamic(a.compthing)),readline.PcItem("get",readline.PcItemDynamic(a.compthing)),readline.PcItem("quit"))
	fmt.Println(ListAllPaths("",maintime))
    l, _ := readline.NewEx(&readline.Config{Prompt: "? ", AutoComplete: pc})
   
	for true {
		command, _ := l.Readline()
		smand := strings.Split(command, " ")
		if smand[0] == "quit" || smand[0] == "quit" {
		    break
		}
           
		if smand[0] == "get" {
		    fmt.Println(SToMap(PathFormat(command[4:]),*maintime).val)
		}
                if smand[0] == "set" {
		    
		    SetValToPath(PathFormat(smand[1]),strings.Join(smand[2:]," "),*maintime)
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
