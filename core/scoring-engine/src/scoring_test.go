package scoring

import (
	"testing"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

func TestScorePrefersLowerLatencyAndHandshake(t *testing.T) {
	s := NewDefault()

	a := contracts.CandidateHealth{LatencyMs: 200, HandshakeOK: true, StabilityScore: 0.8, RecentSuccess: 0.8}
	b := contracts.CandidateHealth{LatencyMs: 900, HandshakeOK: false, StabilityScore: 0.8, RecentSuccess: 0.8}

	if s.Score(a) <= s.Score(b) {
		t.Fatalf("expected a to score higher than b")
	}
}
