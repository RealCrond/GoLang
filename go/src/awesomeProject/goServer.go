package main

import (
	"net/http"
	"fmt"
	"time"
	"html/template"
	"log"
)

func viewData(w http.ResponseWriter,r *http.Request)  {
	//testserverevent.html为html页面的名称
	r.ParseForm()
	if r.Method == "GET"{
		t,err:=template.ParseFiles("testserverevent.html")
		if err != nil{
			fmt.Fprintf(w,"parse template error:%s",err.Error())
			//return
		}
		t.Execute(w,nil)
	}
}

func getDate(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("enter getDate.")
	for{
		w.Header().Set("Content-Type","text/event-stream;charset=utf-8")
		w.Header().Set("Catch-Control","no-cache")
		dtstr :="data:" + fmt.Sprint(time.Now()) +"\n\n"
		dtstr = dtstr + "retry:"+"1000"+"\n\n"
		w.Write([]byte(dtstr))

		if f,ok :=w.(http.Flusher);ok{
			f.Flush()
		}else{
			fmt.Println("no flush")
		}
		time.Sleep(10000*time.Millisecond)
	}
}

func main()  {
	fmt.Println("enter main.")

	http.HandleFunc("/",viewData)
	http.HandleFunc("/getdate",getDate)
	//测试当addr的传入值为空时没法运行的原因
	//err:=http.ListenAndServe("",nil)
	//if err != nil{
	//	log.Fatal("ListenAndServer:",err)
	//}
	//在mac下运行，必须给一个大于1024的端口
	err := http.ListenAndServe(":7000",nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}