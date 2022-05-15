package handlerHttp

import (
	"fmt"
	"github.com/cnblvr/testing_forms/pkg/handler/http/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type Handler struct {
	router *mux.Router
}

func New() *Handler {
	h := new(Handler)
	h.router = mux.NewRouter()

	h.router.Use(middleware.BuildMiddlewareLogger(0))

	h.router.Path("/ping").HandlerFunc(h.Ping).Methods(http.MethodGet)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) PrintHandlers() {
	err := h.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			// is subrouter
			return nil
		}
		if route.GetHandler() == nil {
			log.Warn().Err(fmt.Errorf("path doesn't have handler")).Str("path", path).Send()
			return nil
		}
		log.Debug().Msgf("[%s] %v", strings.Join(methods, ","), path)
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to iterate routes")
	}
}
