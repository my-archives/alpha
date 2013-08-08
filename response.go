package alpha

import (
  "mime"
  "strings"
  "net/http"
)

type Response struct {
  Out       http.ResponseWriter
  Headers   http.Header
}

func (res *Response) Status(code int) *Response {
  res.Out.WriteHeader(code)
  return res
}

func (res *Response) SendByte(body []byte) *Response {
  res.Out.Write(body)
  return res
}

func (res *Response) SendString(body string) *Response {
  res.Out.Write([]byte(body))
  return res
}

/**
 *  Send a response.
 *
 *  Examples:
 *
 *    res.Send("<p>some html</p>")
 *    res.Send([]byte("hello"))
 *    res.Send(404, "Sorry, cant find taht")
 *    res.Send(404)
 *
 */

func (res *Response) Send(args ...interface{}) *Response {
  var (
    c, b interface{}
    l int
    body string
    code = -1
  )

  l = len(args)
  switch {
  case l == 0:
    body = ""
  case l == 1:
    c = args[0]
  case l > 1:
    c = args[0]
    b = args[1]
  }

  switch v := c.(type) {
  case int:
    code = v
    body = http.StatusText(code)
  case string:
    body = v
  case []byte:
    body = string(v)
  }

  if l > 1 {
    switch v := b.(type) {
    case string:
      body = v
    case []byte:
      body = string(v)
    }
  }

  if code != -1 {
    res.Status(code)
  }
  res.SendString(body)
  return res
}

func (res *Response) Type(t string) *Response {
  if t == "" {
    return res
  }

  field := "Content-Type"
  val := strings.ToLower(t)

  if ^strings.Index(val, "/") == 0 {
    if !strings.HasPrefix(val, ".") {
      val = "." + val
    }
    val = mime.TypeByExtension(val)
  }

  return res.SetHeader(field, val)
}

func (res *Response) ContentType(t string) *Response {
  return res.Type(t)
}

func (res *Response) SetHeader(field string, val string) *Response {
  res.Headers.Set(field, val)
  return res
}

func (res *Response) Set() *Response {
  return res
}
