package poll

import routing "github.com/go-ozzo/ozzo-routing/v2"

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/albums/<id>", res.get)
	r.Get("/albums", res.query)

	r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("/albums", res.create)
	r.Put("/albums/<id>", res.update)
	r.Delete("/albums/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}
