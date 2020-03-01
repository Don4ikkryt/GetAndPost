package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", postAndGet)
	http.ListenAndServe(":80", nil)
}

func postAndGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic("sds")
		}
		defer r.Body.Close()
		fmt.Println(string(b))
	case "GET":
		fmt.Println(r.URL.Query())

	default:
		fmt.Println("error")
	}
}
