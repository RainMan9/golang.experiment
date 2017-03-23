package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", HandleRequest)
	http.ListenAndServe(":8888", nil)
}
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>URL参数</h1>"))
	w.Write([]byte(fmt.Sprintf("%v", r.URL.Query())))
}
