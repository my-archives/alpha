package alpha

import (
  "log"
  "net/http"
)

type Address struct {
  Port      string
  Hostname  string
}

type HashObject map[string]interface{}

type Alpha struct {
  Request   *Request
  Response  *Response

  Settings  HashObject
}

func (a *Alpha) init() {
  a.Settings = HashObject{}
}

func (a *Alpha) handle() http.HandlerFunc {
  req := a.Request
  res := a.Response

  return func (w http.ResponseWriter, r *http.Request) {
    req.In = r
    res.Out = w
    req.Res = res
    res.Req = req

    req.Query = r.URL.Query()
    req.Headers = r.Header
    res.Headers = w.Header()

    // test
    res.Charset = "UTF-8"
    res.Type("html")
    res.SetHeader("X-Powered-By", "Alpha")
    res.Send("Hello " + req.Get("User-Agent"));
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
