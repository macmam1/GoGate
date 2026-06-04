package engineadapters

import (
	"errors"
	"testing"
)

func TestMapEngineErrorClassification(t *testing.T) {
	err := MapEngineError("xray", errors.New("permission denied"), "")
	var ae AdapterError
	if !errors.As(err, &ae) {
		t.Fatalf("expected adapter error")
	}
	if ae.Kind != ErrPermission {
		t.Fatalf("expected permission kind, got %s", ae.Kind)
	}
}
