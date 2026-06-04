package engineadapters

import "testing"

func TestSemanticKindFromStderr(t *testing.T) {
	kind := semanticKindFromStderr("failed to bind: address already in use")
	if kind != ErrPortInUse {
		t.Fatalf("expected port in use, got %s", kind)
	}

	kind = semanticKindFromStderr("invalid config: unknown field")
	if kind != ErrConfigInvalid {
		t.Fatalf("expected config invalid, got %s", kind)
	}
}
