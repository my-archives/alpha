package alpha

/**
 * Create an alpha application.
 */

func CreateApplication() *Alpha {
  app := &Alpha{}
  app.Request = &Request{}
  app.init()
  return app
}
