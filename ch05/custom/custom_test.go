package custom

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadContents(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	defer server.Close()

	_, err := ReadContents(server.URL)

	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	httpErr, ok := err.(*HTTPError)
	if !ok {
		t.Fatalf("Expected HTTPError, got %T", err)
	}

	expected := fmt.Sprintf("http status code = %d, url = %s", http.StatusNotFound, server.URL)
	if httpErr.Error() != expected {
		t.Fatalf("Expected error message %q, got %q", expected, httpErr.Error())
	}
}
