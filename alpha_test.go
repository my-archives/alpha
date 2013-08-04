package alpha

import (
  "reflect"
  "fmt"
)

func main () {
  app := CreateApplication()
  a := reflect.ValueOf(app)
  fmt.Println(a.Type())
}
