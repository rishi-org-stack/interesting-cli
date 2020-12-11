package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*Query type*/
type Query struct {
	Query  string
	Effect string
	Path   string
}

/*New function*/
func (q *Query) New(args []string) {
	if len(args) >= 3 {
		q.Query = args[1]
		q.Effect = args[2]
		if len(args) > 3 {
			q.Path = args[3]
		}

	} else {
		fmt.Println("insuffcient number of arguments")
	}

}

/*Makedir function*/
func (q *Query) Makedir() {
	
		err := os.Mkdir(q.Effect, 7)
		if err != nil {
			log.Fatal(err)
		}
	
}

/*Makefile function*/
func (q *Query) Makefile() {
	
	file, err := os.Create(q.Path + q.Effect)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(file)
	}	
}

/*Read function*/
func (q *Query) Read() []byte {
	var b []byte
	con, err := ioutil.ReadFile(q.Path + q.Effect)
	if err != nil {
		log.Fatal(err)
	} else {
		b = con
	}
	return b
}

/*Write function*/
func (q *Query) Write() {
	var s []string
	i := ""
	for {
		var v string
		fmt.Scan(&v)
		s = append(s, v)
		if v == "end" {
			break
		}
	}
	for _, val := range s {
		val = val + " "
		i += val
	}
	mess := []byte(i)
	err := ioutil.WriteFile(q.Path+q.Effect, mess, 7)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Done")
	}
	fmt.Println(s)
}

/*Listdir function*/
func (q *Query) Listdir() {
	if q.Query == "listd" {
		files, err := ioutil.ReadDir(q.Path + q.Effect)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Println("invalid method")
	}
}

/*Getsize funtion*/
func (q *Query) Getsize() {
		files, err := ioutil.ReadDir(q.Path)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if file.Name() == q.Effect {
				if file.Size() <= 1000 {
					fmt.Println(fmt.Sprint(file.Size()) + " B")
				}
				
			}
			
		}
	
}

/*Delete function*/
func (q *Query) Delete() {

		er := os.Remove(q.Path + q.Effect)
		if er != nil {
			log.Fatal(er)
		} else {
			fmt.Println("done")
		}
	
}

/*Move function need correction*/
func (q *Query) Move(args []string) {
	var q1 Query
	q1.Query=q.Query
	q1.Effect=q.Effect
	
	q1.Path=args[3]
	q.Path=""
	q1.Makefile()
	b:=q.Read()
	err:= ioutil.WriteFile(q1.Path+q1.Effect,b,7)
	if err!=nil{
		log.Fatal(err)
	}
	q.Delete()
}

/*Copy function todo commenting*/
func (q *Query)Copy(args []string){
	var q1 Query
	q1.Query=q.Query
	q1.Effect=q.Effect
	
	q1.Path=args[3]
	q.Path=""
	q1.Makefile()
	b:=q.Read()
	err:= ioutil.WriteFile(q1.Path+q1.Effect,b,7)
	if err!=nil{
		log.Fatal(err)
	}
}

/*Find to do commenting*/
func (q*Query)Find(b byte)(bool,int,int){
	mess:= q.Read()
	w:=0
	l:=0
	pres:=false
	for _,val := range mess{
		w++
		if val=='.'{
			l++
		}
		if val==b{
			pres=true
			break
		}
	}
	return pres,w,l
}

//Solver todo commenting
func Solver(){
	var q Query
	args := os.Args
	q.New(args)
	switch args[1] {
	case "md":
		q.Makedir()

	case "mf":
		q.Makefile()

	case"read":
		b:= q.Read()
		fmt.Printf("%s",b)

	case"write":
		q.Write()

	case "listd":
		q.Listdir()

	case "gs":
		q.Getsize()

	case"del":
		q.Delete()

	case "mvf":
		q.Move(args)

	case "cpf":
		q.Copy(args)

	case "find":
		var b string
		fmt.Scan(&b)
		for _,val:= range []byte(b){
			ok,w,l:=q.Find(val)
			fmt.Println(ok)
			fmt.Printf("lines %d\n",l)
			fmt.Printf("words %d\n",w)
		}
	}
}