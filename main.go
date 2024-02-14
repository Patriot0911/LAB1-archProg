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

func (t Timestamp) String() string {
  return t.Value[:19]
}
