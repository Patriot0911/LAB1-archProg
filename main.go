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

func (t Timestamp) String() string {
	return t.Value[:19]
}

func main() {
	fmt.Println("Server is listening on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
