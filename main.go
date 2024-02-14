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