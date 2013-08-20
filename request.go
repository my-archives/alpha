package alpha

import (
  "mime"
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

//
//  Check if the incoming request contains the "Content-Type"
//  header field, and it contains the give mime `type`.
//
//  Examples:
//
//    // With Content-Type: text/html; charset=utf-8
//    req.Is("html")
//    req.Is("text/html")
//    req.Is("text/*")
//    // => true
//

func (req *Request) Is(mtype string) bool {
  ct := req.get("Content-Type")
  if ct == "" {
    return false
  }

  ct = strings.Split(ct, ";")[0]

  if ^strings.Index(mtype, "/") == 0 {
    if !strings.HasPrefix(mtype, ".") {
      mtype = "." + mtype
    }
    mtype = mime.TypeByExtension(mtype)
    mtype = strings.Split(mtype, ";")[0]
  }

  if ^strings.Index(mtype, "*") != 0 {
    ts := strings.Split(mtype, "/")
    cts := strings.Split(ct, "/")
    if "*" == ts[0] && ts[1] == cts[1] {
      return true
    }
    if "*" == ts[1] && ts[0] == cts[0] {
      return true
    }
    return false
  }

  return ^strings.Index(ct, mtype) != 0
}

func (req *Request) Xhr() bool {
  val := req.get("X-Requested-With")
  return val != "" && "xmlhttprequest" == strings.ToLower(val)
}

func (req *Request) Path() string {
  return req.In.URL.Path
}
