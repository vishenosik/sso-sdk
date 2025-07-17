package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	_http "github.com/vishenosik/gocherry/pkg/http"
	"github.com/vishenosik/sso-sdk/errors"
)

var (
	routeSystem = _http.MethodFunc("system")

	PingMethod       = routeSystem("ping")
	GetMetricsMethod = routeSystem("metrics")
)

type Metric struct {
	Param1 string `json:"param_one"`
	Param2 string `json:"param_two"`
	Param3 string `json:"param_three"`
}

type PingResponse struct {
	Message string `json:"message"`
	Search  string `json:"search"`
}

type Metrics = []*Metric

type MetricsUsecase interface {
	GetMetrics() (Metrics, error)
	LogMetrics(Metrics) error
}

type SystemApi struct {
	metrics MetricsUsecase
}

func NewSystemApi(
	metrics MetricsUsecase,
) *SystemApi {
	return &SystemApi{
		metrics: metrics,
	}
}

func (a *SystemApi) Routers(r chi.Router) {

	r.Group(func(r chi.Router) {

		r.With(
			_http.SetHeaders(),
		)

		r.Get(PingMethod, a.Ping())

		r.Route(GetMetricsMethod, func(r chi.Router) {
			r.Get(_http.BlankRoute, a.GetMetrics())
			r.Post("/log", a.LogMetrics())
		})
	})
}

// ping godoc
//
//	@Summary 	Check system health
//	@Tags 		system
//	@Router 	/api/system.ping [get]
//	@Produce 	json
//
//	@Param q query string false "Search query"
//
//	@Success 	200 {object} PingResponse "not ok"
//	@Failure 	406 {string} string    "not ok"
func (a *SystemApi) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		q := r.URL.Query().Get("q")

		if q == "error" {
			http.Error(w, "got error search", http.StatusNotAcceptable)
			return
		}

		response := PingResponse{
			Message: "pong",
			Search:  q,
		}

		resp, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	}
}

func (a *SystemApi) GetMetrics() http.HandlerFunc {

	errmp := httpErrorsMap(map[error]int{
		errors.ErrNotFound: http.StatusNotFound,
	})

	return func(w http.ResponseWriter, r *http.Request) {

		metrics, err := a.metrics.GetMetrics()
		if err != nil {
			http.Error(w, "failed to get metrics", errmp.Get(err))
			return
		}

		if len(metrics) == 0 {
			http.Error(w, "no metrics found", http.StatusNoContent)
			return
		}

		resp, err := json.Marshal(metrics)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

// LogMetrics godoc
//
//	@Summary 	Check system health
//	@Tags 		system
//	@Router 	/api/system.metrics/log [post]
//	@Produce 	json
//
//	@Param metrics body Metrics true "Metrics"
//
//	@Success 	200 {object} PingResponse "not ok"
//	@Failure 	406 {string} string    "not ok"
func (a *SystemApi) LogMetrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metrics, err := _http.Decode[Metrics](r)
		if err != nil {
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		if err := a.metrics.LogMetrics(metrics); err != nil {
			http.Error(w, "failed to log metrics: "+err.Error(), http.StatusBadRequest)
			return
		}

		response := struct {
			Message string `json:"message"`
		}{
			Message: "success",
		}
		resp, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	}
}
