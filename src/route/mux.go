package route

import (
	"net/http"
)

type myHandler struct {
	*http.ServeMux
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultMux.HandleFunc(pattern, handler)
}

func Handle(pattern string, handler http.Handler) {
	DefaultMux.Handle(pattern, handler)
}

var DefaultMux = NewMux()

func NewMux() *myHandler {
	return &myHandler{ServeMux: http.DefaultServeMux}
}

func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.ServeMux.ServeHTTP(w, r)
}
