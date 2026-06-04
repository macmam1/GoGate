package scoring

import (
	"github.com/macmam1/GoGate/core/internal/contracts"
	probe "github.com/macmam1/GoGate/core/probe-engine/src"
)

func HealthFromProbeResult(r probe.Result) contracts.CandidateHealth {
	recent := 0.0
	if r.HandshakeOK && r.Err == nil {
		recent = 1.0
	}
	return contracts.CandidateHealth{
		LatencyMs:      r.LatencyMs,
		PacketLossPct:  r.PacketLossPct,
		HandshakeOK:    r.HandshakeOK && r.Err == nil,
		RecentSuccess:  recent,
		StabilityScore: r.StabilityScore,
	}
}
