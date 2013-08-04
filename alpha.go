package alpha

/**
 * Create an alpha application.
 */

func CreateApplication () *Alpha {
  app := &Alpha{}
  app.init()
  return app
}
