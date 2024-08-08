package main

import (
	"fmt"
	"net/http"
	"time"
)

func getEvents(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")

	for {
		fmt.Fprintf(w, "data: %s\n\n", "Hello world!")
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)
	}
}

func main() {
	http.HandleFunc("/events", getEvents)

	http.ListenAndServe(":8000", nil)
}
