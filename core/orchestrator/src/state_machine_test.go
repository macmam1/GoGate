package orchestrator

import "testing"

func TestStateFlow(t *testing.T) {
	state := StateIdle
	state = NextState(state, EventConnectRequested)
	if state != StateConnecting {
		t.Fatalf("expected connecting, got %s", state)
	}
	state = NextState(state, EventConnectFailed)
	if state != StateFallback {
		t.Fatalf("expected fallback, got %s", state)
	}
	state = NextState(state, EventFallbackNext)
	if state != StateConnecting {
		t.Fatalf("expected connecting after fallback-next, got %s", state)
	}
}
