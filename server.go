package main

import (
	"io"
	"net/http"

	"github.com/golang/glog"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	glog.Info("RUNNING SERVER")
	http.ListenAndServe(":8000", nil)

}
