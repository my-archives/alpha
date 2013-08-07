package alpha

import (
  "log"
  "net/http"
)

type Alpha struct {
  Request   *Request
}

type Address struct {
  Port      string
  Hostname  string
}

func (a *Alpha) init() {}

func (a *Alpha) handle() http.HandlerFunc {
  request := a.Request
  return func (w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    request.In = r
    request.Query = r.Form
    w.Write([]byte("Hello alpha."))
  }
}

func (a *Alpha) Listen(address Address) *http.Server {
  server := &http.Server{
    Addr: address.Hostname + ":" + address.Port,
    Handler: a.handle(),
  }
  log.Fatal(server.ListenAndServe())
  return server
}
