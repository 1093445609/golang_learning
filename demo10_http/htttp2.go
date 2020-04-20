package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 在/后面加上 index ，来指定访问路径
	http.HandleFunc("/index", index2)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
func index2(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("demo10_http/index.html")
	w.Write(content)
}
