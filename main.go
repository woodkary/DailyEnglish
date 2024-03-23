package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./stup.txt")
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hahaha", sayHello)
	err := http.ListenAndServe(":9090", nil) //端口前加冒号，否则listen tcp: address 9090: missing port in address
	if err != nil {
		fmt.Printf("http serve failed, err: %v", err)
		return
	}
}
