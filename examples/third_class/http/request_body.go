package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

}
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body err:%v", err)
		//这里一定要返回，不然还会继续执行
		return
	}
	//需要把[]byte 转换为 string类型
	fmt.Fprintf(w, "read the data :%s \n", string(body))

	//第二次尝试读的时候会发现什么也读不到
	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body err:%v", err)
		//这里一定要返回，不然还会继续执行
		return
	}
	//需要把[]byte 转换为 string类型
	fmt.Fprintf(w, "read the data :%s\n", string(body))

}
func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	//body, _ := r.GetBody()
	//io.ReadAll(body)
	//
	//body, _ = r.GetBody()
	//io.ReadAll(body)

	if r.GetBody == nil {
		fmt.Fprintf(w, "get body is nil")
	} else {
		fmt.Fprintf(w, "get body is not nil")
	}
}
func queryParams(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query()
	//name:=value["name"][0]

	fmt.Fprintf(w, "query is %v\n", value)

}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "headers are %v\n", r.Header)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, "data is %v\n", string(data))
}

func form(w http.ResponseWriter, r *http.Request) {
	//直接调用是没有 form 的
	fmt.Fprintf(w, "before parse form is %v\n", r.Form)
	//使用 r.ParseForm()之后才能使用 r.Form得到 form
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form err:%v", r.Form)
	}
	fmt.Fprintf(w, "after parse form is %v\n", r.Form)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

}
