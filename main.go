package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "time"
)

type Timestamp struct {
	Value string `json:"time"`
}

const (
	PORT = ":8795"
)

func main() {
	fmt.Println("Server is listening on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
