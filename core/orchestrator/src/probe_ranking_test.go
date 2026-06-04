package orchestrator

import (
	"testing"

	"github.com/macmam1/GoGate/core/internal/contracts"
	probe "github.com/macmam1/GoGate/core/probe-engine/src"
	scoring "github.com/macmam1/GoGate/core/scoring-engine/src"
)

func TestRankFromProbeResults(t *testing.T) {
	o := Orchestrator{Scorer: scoring.NewDefault()}
	profiles := []contracts.NormalizedProfile{
		{ID: "a", Protocol: "vless", Endpoint: "a.example", Port: 443},
		{ID: "b", Protocol: "vless", Endpoint: "b.example", Port: 443},
	}
	results := []probe.Result{
		{ProfileID: "b", LatencyMs: 850, HandshakeOK: true, StabilityScore: 0.5},
		{ProfileID: "a", LatencyMs: 120, HandshakeOK: true, StabilityScore: 0.9},
	}
	ranked := o.RankFromProbeResults(profiles, results)
	if len(ranked) != 2 {
		t.Fatalf("expected 2 ranked items, got %d", len(ranked))
	}
	if ranked[0].Profile.ID != "a" {
		t.Fatalf("expected profile a ranked first, got %s", ranked[0].Profile.ID)
	}
}
