package main

import "fmt"
import "github.com/chzyer/readline"
import "regexp"
import "strings"
import "encoding/json"

type mapobj struct {
    Val string
    Sub map[string]*mapobj
}

type completeholder struct {
	M *mapobj 
}

func SToMap(s string,m mapobj) *mapobj {
    r := regexp.MustCompile("^[a-zA-Z0-9]*/")
    a := r.FindString(s)
    path := strings.TrimSuffix(a, "/")
    newa := strings.TrimPrefix(s,a)
    //fmt.Println(path)
    if(len(newa) > 0){
        _,ok := m.Sub[path]
        if(!ok) {
            return MakeMapObj("Item not found")
        }
        return SToMap(newa, *m.Sub[path])
    } else {
	_,ok := m.Sub[path]
        if(!ok){
            return MakeMapObj("Item not found")
        }
        return m.Sub[path]
    }
}

func MakeMapObj(x string) *mapobj {
    ipm := new(mapobj)
    ipm.Val = x
    ipm.Sub = make(map[string]*mapobj)
    return ipm
}


func SetValToPath(s string, val string, m mapobj) *mapobj {
	r := regexp.MustCompile("^[a-zA-Z0-9]*/")
    a := r.FindString(s)
    path := strings.TrimSuffix(a, "/")
    newa := strings.TrimPrefix(s,a)
    //fmt.Println(path)
    if(len(newa) > 0){
        _, ok := m.Sub[path]
		if (!ok) {
			m.Sub[path] = MakeMapObj("sub")
		}
        return SetValToPath(newa, val, *m.Sub[path])
    } else {
		m.Sub[path] = MakeMapObj(val)
		//m.Val = val
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
	for k, v := range m.Sub {
		l = append(l,(prefix + k + "/"))
		if len(v.Sub) > 0 {
			e := ListAllPaths((prefix + k + "/"),v)
			for _,v := range e {
				l = append(l,v)
			}
		}
	}
	return l
}

func (e completeholder ) compthing(a string) []string {
	return ListAllPaths("",e.M)
}

func main(){
    maintime := MakeMapObj("sub")
    maintime.Sub["players"] = MakeMapObj("sub")
    SetValToPath("home/","test",*maintime)
	fmt.Println(SToMap("home/", *maintime).Val)
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
		    fmt.Println(SToMap(PathFormat(command[4:]),*maintime).Val)
		}
                if smand[0] == "set" {
		    
		    SetValToPath(PathFormat(smand[1]),strings.Join(smand[2:]," "),*maintime)
		}
	}
		

	
    readline.SetAutoComplete(pc)
    //b, _ := l.Readline()
    
    //fmt.Println(b)

    main := MakeMapObj("sub")
    main.Sub["hehe"] = MakeMapObj("sub")
    main.Sub["hehe"].Sub["test"] = MakeMapObj("data")
    //fmt.Println(main.Sub["hehe"].Val)
    fmt.Println(SToMap("hehe/",*main).Val)
    b, err := json.Marshal(maintime)
    fmt.Println(err)
    var e mapobj
    fmt.Println(maintime)
    fmt.Println(string(b))
    _ = json.Unmarshal(b, &e)
    fmt.Println(e)
    fmt.Println(e.Sub["home"])
}
