package pprof

import (
	"net/http"
	"net/http/pprof"
	runtime_pprof "runtime/pprof"
	"strings"
)

type handler struct {
}

var profileHandlers map[string]http.Handler

func init() {
	profileHandlers = make(map[string]http.Handler)
	for _, profile := range runtime_pprof.Profiles() {
		profileHandlers[profile.Name()] = pprof.Handler(profile.Name())
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if !strings.HasPrefix(r.URL.Path, "/debug/pprof") {
		next(w, r)
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) > 4 {
		next(w, r)
		return
	}

	if len(parts) == 3 || (len(parts) == 4 && parts[3] == "") {
		pprof.Index(w, r)
		return
	}

	handler := profileHandlers[parts[3]]
	if handler == nil {
		next(w, r)
		return
	}
	handler.ServeHTTP(w, r)
	return
}

// Pprof returns a handler which will serve pprof data for the path /debug/pprof
func Pprof() *handler {
	return &handler{}
}
