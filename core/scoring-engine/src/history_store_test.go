package scoring

import (
	"path/filepath"
	"testing"
)

func TestMemoryHistoryStore(t *testing.T) {
	s := NewMemoryHistoryStore()
	_ = s.Record("p1", true)
	_ = s.Record("p1", false)
	v, ok := s.RecentSuccess("p1")
	if !ok {
		t.Fatalf("expected history present")
	}
	if v <= 0 || v >= 1 {
		t.Fatalf("expected ratio between 0 and 1, got %f", v)
	}
}

func TestFileHistoryStore(t *testing.T) {
	s := &FileHistoryStore{Path: filepath.Join(t.TempDir(), "history", "scores.json")}
	if err := s.Record("p1", true); err != nil {
		t.Fatalf("record: %v", err)
	}
	if err := s.Record("p1", false); err != nil {
		t.Fatalf("record: %v", err)
	}
	v, ok := s.RecentSuccess("p1")
	if !ok {
		t.Fatalf("expected history present")
	}
	if v <= 0 || v >= 1 {
		t.Fatalf("expected ratio between 0 and 1, got %f", v)
	}
}
