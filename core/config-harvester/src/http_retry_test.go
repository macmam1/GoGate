package harvester

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestFetchSubscriptionWithRetry(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = fmt.Fprintln(w, "temporary")
			return
		}
		_, _ = fmt.Fprintln(w, "# comment")
		_, _ = fmt.Fprintln(w, "id1:vless:example.com")
	}))
	defer srv.Close()

	entries, err := FetchSubscriptionWithClient(context.Background(), srv.URL, defaultHTTPClient())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 parsed entry, got %d", len(entries))
	}
}
