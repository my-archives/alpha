package alpha

import (
  "log"
  "net/http"
)

type Alpha struct {}

func (a *Alpha) init() {}

func (a *Alpha) Listen(addr string) *http.Server {
  server := &http.Server{
    Addr: addr,
  }

  log.Fatal(server.ListenAndServe())
  return server
}
