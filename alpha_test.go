package alpha

import (
  "fmt"
  "os"
  "reflect"
)


func main() {
  var address = Address{}
  address.Port = os.Getenv("PORT")
  if address.Port == "" {
    address.Port = "3000"
  }

  app := CreateApplication()
  app.Listen(address)
  a := reflect.ValueOf(app)
  fmt.Println(a.Type())
}
