// Send an infinite stream of unclosed HTML tags.

package main

import (
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
	tags := []string{
		"a", "b", "code", "div", "em", "fieldset", "form", "h1", "h2", "h3",
		"i", "input", "label", "legend", "p", "pre", "q", "s", "small", "span",
		"strong", "sub", "sup", "table", "tbody", "td", "tfoot", "th", "thead",
		"tr", "u",
	}
	for {
		for i := 0; i < 1000; i++ {
			code := "<" + tags[rand.Intn(len(tags))] + ">"
			_, err := w.Write([]byte(code))
			if err != nil {
				return
			}
		}
		flusher.Flush()
	}
}
