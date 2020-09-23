package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetuser(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://127.0.0.1:8000/123", nil)
	GetUser(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("the mapper is not working good %s", "$")
	}
}
