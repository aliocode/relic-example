package httpserver

import (
	"net/http"

	api "github.com/aliocode/relic-example/api/http"
	"github.com/go-chi/chi/v5"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Server struct{}

func New(handler api.ServerInterface, relic *newrelic.Application) http.Handler {
	r := chi.NewRouter()
	r.Use(relicMiddleware(relic))
	r.Get("/healthcheck", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNoContent)
	})
	return api.HandlerFromMux(handler, r)
}

func relicMiddleware(app *newrelic.Application) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			txn := app.StartTransaction(r.Method + r.URL.RequestURI())
			defer txn.End()
			txn.SetWebRequestHTTP(r)
			w = txn.SetWebResponse(w)
			r = newrelic.RequestWithTransactionContext(r, txn)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
