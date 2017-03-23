package main

import (
	"net/http"
)

func main() {
	handler := &HttpHandler{}
	http.ListenAndServe(":8888", handler)
}

type HttpHandler struct {
}

func (this *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>在ServerHTTP里</h1>"))
	w.Write([]byte(r.URL.Path))
}
