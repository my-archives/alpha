package alpha

import (
  "reflect"
  "fmt"
)

func main () {
  app := CreateApplication()
  app.Listen(":3000")
  a := reflect.ValueOf(app)
  fmt.Println(a.Type())
}
