package query

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	dt "github.com/rishi-org-stack/cli/events/data"

	ut "github.com/rishi-org-stack/cli/events/utils"

	mn "github.com/rishi-org-stack/cli/events/manager"
	"time"
)

type WhattoDo struct {
	What  string
	Which interface{}
}

var d = "cli"
var c = "event"



func Isvalid(s string)bool{
	ans :=false
	dm:=make(map[int]int)
	dm[1] =31
	dm[2] =28
	dm[3] =31
	dm[4] =30
	dm[5] =31
	dm[6] =30
	dm[7] =31
	dm[8] =31
	dm[9] =30
	dm[10]=31
	dm[11]=30
	dm[12]=31
	// year,m,d:=time.Now().Date()
	gd,gm:=ut.Parse(s)
	if gm>=1&&gm<13{

		if gd<=dm[gm]&&gd>0{
			ans =true
		}else{
			ans = false
			fmt.Println("Nt a valid date for given month")
		}
	}else{
		ans = false
		fmt.Println("invalid month")
	}
	return ans 
}

func (w *WhattoDo) List(s1,s2 string) {
var e dt.Event
	var data []bson.M
	w.What = s1
	w.Which = s2
	data = e.GetAll(d, c)
	for _,dat:= range data{
		fmt.Println(dat["name"].(string))
		fmt.Println(dat["date"].(string))
		fmt.Println(dat["des"].(string))
		fmt.Println()
	}
}


func (w *WhattoDo)GetoneindetailByname(s1,s2 string){
var e dt.Event
	_,m,day:=time.Now().Date()
	//impplent algorithm for time left
	var details bson.M
	w.What = s1
	w.Which = s2
	e.Name= w.Which.(string)
	details = e.Getone(d,c)
	if Isvalid(details["date"].(string)){
		gd,gm:=ut.Parse(details["date"].(string))
		fmt.Printf("no of days left->%d\n",gd-day)
		fmt.Printf("no of months left->%d\n",gm-int(m))
	}else{
		fmt.Println("what")
	}
}

func (w *WhattoDo)GetoneBydate(s1,s2 string){
var e dt.Event
var data []bson.M
	_,m,day:=time.Now().Date()
	w.What= s1
	w.Which=s2
	e.Date=w.Which.(string)
	data= e.GetAll(d,c)
	for _,val:= range data{
		if e.Date==val["date"].(string){
			if Isvalid(val["date"].(string)){
				gd,gm:=ut.Parse(val["date"].(string))
				fmt.Printf("no of days left->%d\n",gd-day)
				fmt.Printf("no of months left->%d\n",gm-int(m))
			}else{
		
			}
		}
	}
}


func (w *WhattoDo)Updateone(s string,id int){
	var e dt.Event
	var s1,s2 string
	w.Which=id
	w.What=s
	e.UID= w.Which.(int32)
	fmt.Println("new name of evemt pls:->")
	fmt.Scan(&s1)
	e.Name=s1
	fmt.Println("New date pls:->")
	fmt.Scan(&s2)
	if Isvalid(s2)==true{
		e.Date = s2
		e.Update(d,c)
	}
	
}


func (w *WhattoDo)Add(s1,s2 string){
	var e dt.Event
	var uid int32
	var des ,date string
	w.What = s1
	w.Which = s2
	e.Name = w.Which.(string)
	fmt.Scan(&uid)
	e.UID = uid
	fmt.Scan(&date)
	if Isvalid(date){
		e.Date = date
		fmt.Scan(&des)
		e.Description = des
		e.Insert(d,c)
	}else{
		fmt.Println("not a valid date")
	}	
}
func (w*WhattoDo)Delete(s string,id int)  {
	var e dt.Event
	w.What=s
	w.Which=id
	e.UID= w.Which.(int32)
	e.Delete(d,c)
}

func Solverquery(){
	var w WhattoDo
	var s1,s2 string
	for {
		fmt.Printf("->\t")
		fmt.Scan(&s1)
		switch s1 {
		case "list":
			fmt.Scan(&s2)
			mn.Olddates()
			w.List(s1,s2)
		case "gon":
			fmt.Scan(&s2)
			w.GetoneindetailByname(s1,s2)
		case "god":
			fmt.Scan(&s2)
			w.GetoneBydate(s1,s2)
		case "upo":
			var id int
			fmt.Scan(&id)
			w.Updateone(s1,id)
		case "add":
			fmt.Println("name pls")
			fmt.Scan(&s2)
			w.Add(s1,s2)
		case "del":
			fmt.Println("id pls")
			var id int
			fmt.Scan(&id)
			w.Delete(s1,id)
		}
		if s1=="stop"{
			fmt.Println("bye")
			break
		}
	}
}