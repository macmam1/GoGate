package scoring

import (
	"testing"

	probe "github.com/macmam1/GoGate/core/probe-engine/src"
)

func TestHealthFromProbeResult(t *testing.T) {
	h := HealthFromProbeResult(probe.Result{LatencyMs: 120, HandshakeOK: true, StabilityScore: 0.7})
	if !h.HandshakeOK {
		t.Fatalf("expected handshake ok")
	}
	if h.RecentSuccess <= 0 {
		t.Fatalf("expected recent success > 0")
	}
}
