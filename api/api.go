package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/vishenosik/gocherry/pkg/errors"
	_http "github.com/vishenosik/gocherry/pkg/http"

	_ "github.com/vishenosik/sso-sdk/gen/swagger"
	"google.golang.org/grpc/codes"
)

type Service interface {
	Routers(r chi.Router)
}

// @title           sso
// @version         0.0.1
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
//
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host      localhost:8080
// @BasePath  /api
// @schemes http https
//
// @securityDefinitions.basic  BasicAuth
//
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func NewHttpHandler(services ...Service) http.Handler {

	router := chi.NewRouter()

	router.Use(
		_http.RequestLogger(),
	)

	router.Route("/api", func(r chi.Router) {

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("doc.json"),
		))

		for i := range services {
			services[i].Routers(r)
		}

	})

	return router
}

type httpErrMap interface {
	Get(err error) int
}

type grpcErrMap interface {
	Get(err error) codes.Code
}

func httpErrorsMap(_map_ map[error]int) httpErrMap {
	return errors.NewErrorsMap(http.StatusInternalServerError, _map_)
}

func grpcErrorsMap(_map_ map[error]codes.Code) grpcErrMap {
	return errors.NewErrorsMap(codes.Internal, _map_)
}
