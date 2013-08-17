package alpha

import (
  "strings"
  "net/url"
  "net/http"
)

type Request struct {
  In        *http.Request
  Res       *Response
  Headers   http.Header
  Query     url.Values
}

func (req *Request) get(field string) string {
  var val string

  field = strings.Title(strings.ToLower(field))

  if field == "Referer" || field == "Referrer" {
    val = req.Headers.Get("Referrer")
    if val == "" {
      val = req.Headers.Get("Referer")
    }
  } else {
    val = req.Headers.Get(field)
  }

  return val
}

func (req *Request) Get(field string) string {
  return req.get(field)
}

func (req *Request) Header(field string) string {
  return req.get(field)
}
