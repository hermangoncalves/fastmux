package fastmux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         Params
}

func (ctx *Context) JSON(code int, data any) {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(code)
	json.NewEncoder(ctx.ResponseWriter).Encode(data)
}

func (ctx *Context) Param(key string) (string, bool) {
	for _, p := range ctx.Params {
		if p.Key == key {
			return p.Value, true
		}
	}

	return "", false
}

type HandlerFunc func(ctx *Context)

// Param represents a single route parameter
type Param struct {
	Key   string
	Value string
}

type Params []Param

// route represents a single route
type route struct {
	method  string
	pattern string
	handler HandlerFunc
}

// Fastmux is the main Fastmux structure
type Fastmux struct {
	routes   []route
	notFound HandlerFunc
}

type H map[string]any

// New creates a new Fastmux
func New() *Fastmux {
	return &Fastmux{
		routes: make([]route, 0),
		notFound: func(ctx *Context) {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": "page not found",
			})
		},
	}
}

// Handle registers a handler for a specific method and pattern
func (mx *Fastmux) Handle(method, pattern string, handler HandlerFunc) {
	if method == "" {
		panic("method must not be empty")
	}

	if len(pattern) < 1 || pattern[0] != '/' {
		panic("path must begin with '/' in path '" + pattern + "'")
	}

	if handler == nil {
		panic("handler must not be nil")
	}

	mx.routes = append(mx.routes, route{
		method:  strings.ToUpper(method),
		pattern: pattern,
		handler: handler,
	})
}

func (mx *Fastmux) Handler() http.Handler {
	return mx
}

// HTTP method-specific handlers
func (mx *Fastmux) GET(pattern string, handler HandlerFunc) {
	mx.Handle(http.MethodGet, pattern, handler)
}

func (mx *Fastmux) POST(pattern string, handler HandlerFunc) {
	mx.Handle(http.MethodPost, pattern, handler)
}

func (mx *Fastmux) PUT(pattern string, handler HandlerFunc) {
	mx.Handle(http.MethodPut, pattern, handler)
}

func (mx *Fastmux) PATCH(pattern string, handler HandlerFunc) {
	mx.Handle(http.MethodPatch, pattern, handler)
}

func (mx *Fastmux) DELETE(pattern string, handler HandlerFunc) {
	mx.Handle(http.MethodDelete, pattern, handler)
}

func (mx *Fastmux) HEAD(pattern string, handler HandlerFunc) {
	mx.Handle(http.MethodHead, pattern, handler)
}

func (mx *Fastmux) OPTIONS(pattern string, handler HandlerFunc) {
	mx.Handle(http.MethodOptions, pattern, handler)
}

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
func (mx *Fastmux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqMethod := req.Method
	reqPath := req.URL.Path

	for _, route := range mx.routes {
		if route.method != reqMethod {
			continue
		}
		if matched, params := matchRoute(route.pattern, reqPath); matched {
			ctx := &Context{ResponseWriter: w, Request: req, Params: params}
			route.handler(ctx)
			return
		}
	}

	ctx := &Context{ResponseWriter: w, Request: req, Params: nil}
	mx.notFound(ctx)
}

func (mx *Fastmux) Run(addr string) error {
	address := resolveAddress(addr)
	debugPrint("Server running on %s", address)
	return http.ListenAndServe(address, mx.Handler())
}

func resolveAddress(addr string) string {
	switch addr {
	case "":
		if port := os.Getenv("PORT"); port != "" {
			return fmt.Sprintf(":%s", port)
		}
		return ":8081"
	default:
		return addr
	}
}
