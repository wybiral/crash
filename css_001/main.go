// Send an infinite stream of garbage as a CSS file.

package main

import (
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/asset.css", asset)
	http.HandleFunc("/", index)
	log.Println("Serving on :8080")
	http.ListenAndServe(":8080", nil)
}

func asset(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	buffer := make([]byte, 1024)
	for {
		rand.Read(buffer)
		_, err := w.Write(buffer)
		if err != nil {
			return
		}
		flusher.Flush()
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<html>
	<head>
		<link rel="stylesheet" href="asset.css">
	</head>
	<body>
	</body>
</html>`))
}
