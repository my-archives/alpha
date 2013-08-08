alpha
-----

```go
package main

import (
  "os"
  α github.com/cfddream/alpha"
)

func main() {
  var address = α.Address{}
  address.Port = os.Getenv("PORT")
  if address.Port == "" {
    address.Port = "3000"
  }

  app := α.CreateApplication()
  app.Listen(address)
}
```
