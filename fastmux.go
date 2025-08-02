package main

import (
	"net/http"
	"strings"
)

type route struct {
	method  string
	pattern string
	handler http.Handler
}

type Router struct {
	routes   []route
	notFound http.Handler
}


func New() *Router {
	return &Router{
		routes:   make([]route, 0),
		notFound: http.NotFoundHandler(),
	}
}

func (r *Router) Handle(method, pattern string, handler http.Handler) {
	r.routes = append(r.routes, route{
		method:  strings.ToUpper(method),
		pattern: pattern,
		handler: handler,
	})
}

func (r *Router) HandleFunc(method, pattern string, handlerFunc http.HandlerFunc) {
	r.Handle(method, pattern, handlerFunc)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, rt := range r.routes {
		if rt.method == req.Method && rt.pattern == req.URL.Path {
			rt.handler.ServeHTTP(w, req)
			return
		}
	}

	r.notFound.ServeHTTP(w, req)
}

func (r *Router) NotFound(handler http.Handler) {
	r.notFound = handler
}
