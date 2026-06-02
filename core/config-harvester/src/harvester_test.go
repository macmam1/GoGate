package harvester

import "testing"

func TestNormalizeAndDedupe(t *testing.T) {
	raw := []RawEntry{
		{Source: "srcA", Value: "id1:vless:example.com"},
		{Source: "srcA", Value: "id1:vless:example.com"},
		{Source: "srcA", Value: "id2:vmess:example.net"},
	}
	profiles, err := Normalize(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(profiles) != 2 {
		t.Fatalf("expected 2 unique profiles, got %d", len(profiles))
	}
}
