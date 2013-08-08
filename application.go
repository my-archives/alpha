package alpha

import (
  "log"
  "net/http"
)

type Address struct {
  Port      string
  Hostname  string
}

type Alpha struct {
  Request   *Request
  Response  *Response
}

func (a *Alpha) init() {}

func (a *Alpha) handle() http.HandlerFunc {
  req := a.Request
  res := a.Response

  return func (w http.ResponseWriter, r *http.Request) {
    req.In = r
    res.Out = w

    req.Query = r.URL.Query()
    res.Headers = w.Header()

    // test
    res.Type("html")
    res.SetHeader("X-Powered-By", "Alpha")
    res.Send("Hello Web!");
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
