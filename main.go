package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	PORT = ":8795"
)

type Timestamp struct {
	Value string `json:"time"`
}

func (t Timestamp) String() string {
	return t.Value[:19]
}

func init() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		timestamp := Timestamp{time.Now().Format(time.RFC3339)}
		defer fmt.Printf("%v\t%s\t%s\n", timestamp, r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timestamp)
	})
}

func main() {
	fmt.Println("Server is listenning on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
