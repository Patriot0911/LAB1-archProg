package main

import (
	"fmt"
	"log"
	"net/http"
)

type Timestamp struct {
	Value string `json:"time"`
}

const (
	PORT = ":8795"
)

func init() {
  http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
    timestamp := Timestamp{time.Now().Format(time.RFC3339)}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(timestamp)
  })
}

func (t Timestamp) String() string {
	return t.Value[:19]
}

/// TTOTOTOTOTOT

func main() {
	fmt.Println("Server is listening on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
