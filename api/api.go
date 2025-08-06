package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/vishenosik/gocherry/pkg/errors"
	_http "github.com/vishenosik/gocherry/pkg/http"
	"github.com/vishenosik/gocherry/pkg/httpSwagger"

	"github.com/vishenosik/sso-sdk/gen/swagger"
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
// @servers.url http://localhost:8080
//
// @securityDefinitions.basic  BasicAuth
//
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func NewHttpHandler(services ...Service) http.Handler {

	services = append(services, []Service{
		httpSwagger.NewSwagger(swagger.SwaggerInfo),
	}...)

	r := chi.NewRouter()
	r.Use(
		middleware.Recoverer,
		_http.RequestLogger(),
	)

	r.Route("/api", func(r chi.Router) {
		for i := range services {
			services[i].Routers(r)
		}
	})

	return r
}

type grpcErrMap interface {
	Get(err error) codes.Code
}

func grpcErrorsMap(_map_ map[error]codes.Code) grpcErrMap {
	return errors.NewErrorsMap(codes.Internal, _map_)
}
