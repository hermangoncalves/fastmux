package fastmux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterParam(t *testing.T) {
	router := New()

	router.Handle(http.MethodGet, "/hello/:name", func(ctx *Context) {
		got, _ := ctx.Param("name")
		want := "world"

		if got != want {
			t.Fatalf("wrong: want %s, got %s", want, got)
		}

		ctx.JSON(http.StatusOK, H{
			"hello": got,
		})
	})

	req, err := http.NewRequest(http.MethodGet, "/hello/world", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected status: got %d, want %d", w.Code, http.StatusOK)
	}
}

func TestRoutes(t *testing.T) {
	var get, head, options, post, put, patch, delete bool

	router := New()

	router.GET("/GET", func(ctx *Context) {
		get = true
	})

	router.HEAD("/GET", func(ctx *Context) {
		head = true
	})

	router.OPTIONS("/GET", func(ctx *Context) {
		options = true
	})

	router.POST("/POST", func(ctx *Context) {
		post = true
	})

	router.PUT("/PUT", func(ctx *Context) {
		put = true
	})

	router.PATCH("/PATCH", func(ctx *Context) {
		patch = true
	})

	router.DELETE("/DELETE", func(ctx *Context) {
		delete = true
	})

	w := httptest.NewRecorder()

	r, _ := http.NewRequest(http.MethodGet, "/GET", nil)
	router.ServeHTTP(w, r)
	if !get {
		t.Error("routing GET failed")
	}

	r, _ = http.NewRequest(http.MethodHead, "/GET", nil)
	router.ServeHTTP(w, r)
	if !head {
		t.Error("routing HEAD failed")
	}

	r, _ = http.NewRequest(http.MethodOptions, "/GET", nil)
	router.ServeHTTP(w, r)
	if !options {
		t.Error("routing OPTIONS failed")
	}

	r, _ = http.NewRequest(http.MethodPost, "/POST", nil)
	router.ServeHTTP(w, r)
	if !post {
		t.Error("routing POST failed")
	}

	r, _ = http.NewRequest(http.MethodPut, "/PUT", nil)
	router.ServeHTTP(w, r)
	if !put {
		t.Error("routing PUT failed")
	}

	r, _ = http.NewRequest(http.MethodPatch, "/PATCH", nil)
	router.ServeHTTP(w, r)
	if !patch {
		t.Error("routing PATCH failed")
	}

	r, _ = http.NewRequest(http.MethodDelete, "/DELETE", nil)
	router.ServeHTTP(w, r)
	if !delete {
		t.Error("routing DELETE failed")
	}
}
