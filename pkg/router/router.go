package router

import (
	"github.com/arionalmond/go-api-boilerplate/pkg/datastore"
	"github.com/arionalmond/go-api-boilerplate/pkg/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// GetRouter returns a chi.Router to power the api
func GetRouter(ds *datastore.MySQLDS, appEnv string) chi.Router {
	// monitor := GetMonitoring()

	rtr := chi.NewRouter()

	rtr.Use(middleware.RequestID)
	rtr.Use(middleware.RealIP)
	rtr.Use(middleware.Recoverer)

	// if appEnv == "PRODUCTION" || appEnv == "STAGE" {
	//  // auth middleware
	// 	rtr.Use()
	// }

	rtr.Group(func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {
			r.Get("/employees/{id}", handler.GetEmployees(ds))
		})

	})

	return rtr
}
