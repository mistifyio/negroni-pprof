package pprof_test

import (
	"fmt"
	"github.com/mistifyio/negroni-pprof"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testHTTPContent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test content")
}

func Test_Index(t *testing.T) {
	w := httptest.NewRecorder()
	handler := pprof.Pprof()

	req, err := http.NewRequest("GET", "http://localhost/debug/pprof", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(w, req, testHTTPContent)

	if !strings.Contains(w.Body.String(), "full goroutine stack dump") {
		t.Fail()
	}
}

func Test_Heap(t *testing.T) {
	w := httptest.NewRecorder()
	handler := pprof.Pprof()

	req, err := http.NewRequest("GET", "http://localhost/debug/pprof/heap", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(w, req, testHTTPContent)

	if !strings.Contains(w.Body.String(), "heap profile:") {
		t.Fail()
	}
}
