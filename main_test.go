package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux("v-test", "h1").ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if !strings.Contains(rec.Body.String(), "labapp v-test running on h1") {
		t.Fatalf("首页内容不对: %q", rec.Body.String())
	}
}

func TestHealthz(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux("v", "h").ServeHTTP(rec, httptest.NewRequest("GET", "/healthz", nil))
	if rec.Code != 200 {
		t.Fatalf("/healthz 返回 %d，应该是 200", rec.Code)
	}
}