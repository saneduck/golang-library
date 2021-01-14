package handlers

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

type Hello struct {
  l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
  return &Hello { l }
}

func (h *Hello) ServeHTTP(rw *http.ResponseWriter, r *http.Request) {
  log.Println("Hello => ServeHTTP")
  d, err := ioutil.ReadAll(r.Body)
  if err != nil {
    http.Error(rw, "Oops", http.StatusBadRequest)
    return
  }
  
  fmt.Fprintf(rw, "Hello %s", d)
}
