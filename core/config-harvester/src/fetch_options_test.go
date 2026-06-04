package harvester

import "testing"

func TestNormalizeFetchOptionsDefaults(t *testing.T) {
	cfg := normalizeFetchOptions(FetchOptions{})
	if cfg.Retry.Attempts == 0 {
		t.Fatalf("expected default retry attempts")
	}
	if cfg.UserAgent == "" {
		t.Fatalf("expected default user-agent")
	}
}
