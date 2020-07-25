package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chiDispatcher = chi.NewRouter()
)

type chiRouter struct{}

// NewChiRouter creates a new Chi router
func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVER(port string) {
	fmt.Printf("Chi HTTP server runing on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
