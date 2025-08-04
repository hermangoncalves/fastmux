package fastmux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
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
