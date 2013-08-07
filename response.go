package alpha

import (
  "log"
  "io"
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

func (res *Response) SendString(body string) *Response {
  if _, err := io.WriteString(res.Out, body); err != nil {
    log.Fatal(err)
  }
  return res
}

func (res *Response) Send() *Response {
  return res
}

func (res *Response) SetHeader(field string, val string) *Response {
  res.Headers.Set(field, val)
  return res
}

func (res *Response) Set() *Response {
  return res
}
