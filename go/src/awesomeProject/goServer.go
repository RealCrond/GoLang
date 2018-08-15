package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func viewData(w http.ResponseWriter, r *http.Request) {
	//testserverevent.html为html页面的名称
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("testserverevent.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error:%s", err.Error())
			//return
		}
		t.Execute(w, nil)
	}
}

func getDate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter getDate.")
	for {
		w.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
		w.Header().Set("Catch-Control", "no-cache")
		dtstr := "data:" + fmt.Sprint(time.Now()) + "\n\n"
		//dtstr = dtstr + "retry:" + "1000" + "\n\n"
		fmt.Println(dtstr)
		w.Write([]byte(dtstr))
		fmt.Println("getDate()")

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			fmt.Println("f.flush")
		} else {
			fmt.Println("no flush")
		}
		time.Sleep(10000 * time.Millisecond)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error:%s", err.Error())
			//return
		}
		t.Execute(w, nil)
	}
}

func initHandleFunc(){
	http.HandleFunc("/", viewData)
	http.HandleFunc("/login", login)
	http.HandleFunc("/getdate", getDate)
}

func main() {
	fmt.Println("enter main.")

	initHandleFunc()

	//在mac下运行，必须给一个大于1024的端口
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
