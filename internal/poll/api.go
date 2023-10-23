package poll

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	//res := resource{service,
	// register endpoints related to albums here
}

type resource struct {
	service Service
	logger  log.Logger
}
