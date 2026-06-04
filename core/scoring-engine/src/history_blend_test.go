package scoring

import "testing"

type fakeHistory struct{}

func (fakeHistory) RecentSuccess(profileID string) (float64, bool) {
	if profileID == "p1" {
		return 0.2, true
	}
	return 0, false
}

func TestBlendScore(t *testing.T) {
	cfg := DefaultBlendConfig()
	out := BlendScore(0.9, "p1", fakeHistory{}, cfg)
	if out >= 0.9 {
		t.Fatalf("expected blended score lower than current because history is poor")
	}

	unchanged := BlendScore(0.7, "unknown", fakeHistory{}, cfg)
	if unchanged != 0.7 {
		t.Fatalf("expected unchanged score for unknown profile")
	}
}
