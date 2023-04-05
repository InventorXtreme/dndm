package main

import "fmt"
import "github.com/chzyer/readline"
import "regexp"
import "strings"
import "encoding/json"
import "io/ioutil"
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

func SaveMapObjToFile(name string, m *mapobj) {
	jsonstring, _ := json.Marshal(m)
	_ = ioutil.WriteFile(name, jsonstring, 0644)
	
}

func ReadMapObjFromFile(name string, m *mapobj) {
	filecontent, _ := ioutil.ReadFile(name)
	_ = json.Unmarshal(filecontent,m)
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
			if len(smand) > 1 && len(smand[1]) > 0 {
				fmt.Println(SToMap(PathFormat(smand[1]),*maintime).Val)
			} else {
				fmt.Println("ERR MUST SUPPLY PATH")
			}
		}
                if smand[0] == "set" {
			if len(smand) > 1 && len(smand[1]) > 0 {
				SetValToPath(PathFormat(smand[1]),strings.Join(smand[2:]," "),*maintime)
			} else {
				fmt.Println("ERR MUST SUPPLY PATH")
			} 
		}
		if smand[0] == "save" {
			fname := "default.json"
			if len(smand) > 1 {
				if len(smand[1]) > 0 {
					fname = smand[1]
				}
			}
			SaveMapObjToFile(fname,maintime)
		}
		if smand[0] == "load" {
			fname := "default.json"
			if len(smand) > 1 {
				if len(smand[1]) > 0 {
					fname = smand[1]
				}
			}
			ReadMapObjFromFile(fname,maintime)
		}
	}
		

	
	readline.SetAutoComplete(pc)
	//b, _ := l.Readline()
    
	//fmt.Println(b)


}
