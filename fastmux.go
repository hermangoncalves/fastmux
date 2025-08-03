package fastmux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         []Param
}

func (ctx *Context) JSON(code int, data any) {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(code)
	json.NewEncoder(ctx.ResponseWriter).Encode(data)
}

func (ctx *Context) Param(key string) string {
	for _, p := range ctx.Params {
		if p.Key == key {
			return p.Value
		}
	}
	return ""
}

type HandlerFunc func(ctx *Context)

// Param represents a single route parameter
type Param struct {
	Key   string
	Value string
}

// route represents a single route
type route struct {
	method  string
	pattern string
	handler HandlerFunc
}

// Fastmux is the main Fastmux structure
type Fastmux struct {
	routes   []route
	notFound http.Handler
}

// New creates a new Fastmux
func New() *Fastmux {
	return &Fastmux{
		routes:   make([]route, 0),
		notFound: http.NotFoundHandler(),
	}
}

// Handle registers a handler for a specific method and pattern
func (r *Fastmux) Handle(method, pattern string, handler HandlerFunc) {
	r.routes = append(r.routes, route{
		method:  strings.ToUpper(method),
		pattern: pattern,
		handler: handler,
	})
}

// HTTP method-specific handlers
func (r *Fastmux) GET(pattern string, handler HandlerFunc) {
	r.Handle(http.MethodGet, pattern, handler)
}

// func (r *Fastmux) POST(pattern string, handle Handle) {
// 	r.Handle(http.MethodPost, pattern, handle)
// }

// func (r *Fastmux) PUT(pattern string, handle Handle) {
// 	r.Handle(http.MethodPut, pattern, handle)
// }

// func (r *Fastmux) PATCH(pattern string, handle Handle) {
// 	r.Handle(http.MethodPatch, pattern, handle)
// }

// func (r *Fastmux) DELETE(pattern string, handle Handle) {
// 	r.Handle(http.MethodDelete, pattern, handle)
// }

// func (r *Fastmux) HEAD(pattern string, handle Handle) {
// 	r.Handle(http.MethodHead, pattern, handle)
// }

// func (r *Fastmux) OPTIONS(pattern string, handle Handle) {
// 	r.Handle(http.MethodOptions, pattern, handle)
// }

// matchRoute checks if a pattern matches a path and extracts parameters
func matchRoute(pattern, path string) (bool, []Param) {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(patternParts) != len(pathParts) {
		return false, nil
	}

	params := make([]Param, 0)
	for i := 0; i < len(patternParts); i++ {
		pp := patternParts[i]
		cp := pathParts[i]
		if strings.HasPrefix(pp, ":") {
			params = append(params, Param{
				Key:   pp[1:],
				Value: cp,
			})
		} else if pp != cp {
			return false, nil
		}
	}
	return true, params
}

// ServeHTTP implements the http.Handler interface
func (r *Fastmux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqMethod := req.Method
	reqPath := req.URL.Path

	fmt.Println(reqMethod)
	fmt.Println(reqPath)

	for _, route := range r.routes {
		if route.method != reqMethod {
			continue
		}
		if matched, params := matchRoute(route.pattern, reqPath); matched {
			ctx := &Context{ResponseWriter: w, Request: req, Params: params}
			route.handler(ctx)
			return
		}
	}
	r.notFound.ServeHTTP(w, req)
}

// NotFound sets a custom 404 handler
func (r *Fastmux) NotFound(handler http.Handler) {
	r.notFound = handler
}
