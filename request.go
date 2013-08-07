package alpha

import (
  "net/url"
  "net/http"
)

type Request struct {
  In    *http.Request
  Query url.Values
}
