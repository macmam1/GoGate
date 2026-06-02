package scoring

import "github.com/macmam1/GoGate/core/internal/contracts"

type Weights struct {
	Latency         float64
	Stability       float64
	Handshake       float64
	RecentSuccess   float64
}

type Scorer struct {
	W Weights
}

func NewDefault() Scorer {
	return Scorer{W: Weights{Latency: 0.35, Stability: 0.35, Handshake: 0.20, RecentSuccess: 0.10}}
}

func (s Scorer) Score(h contracts.CandidateHealth) float64 {
	latencyComponent := 1.0 - minFloat(float64(h.LatencyMs)/1200.0, 1.0)
	handshake := 0.0
	if h.HandshakeOK {
		handshake = 1.0
	}
	return s.W.Latency*latencyComponent +
		s.W.Stability*clamp(h.StabilityScore) +
		s.W.Handshake*handshake +
		s.W.RecentSuccess*clamp(h.RecentSuccess)
}

func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
