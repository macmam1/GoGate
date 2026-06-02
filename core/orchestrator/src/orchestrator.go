package orchestrator

import (
	"sort"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type CandidateScorer interface {
	Score(h contracts.CandidateHealth) float64
}

type CandidateInput struct {
	Profile contracts.NormalizedProfile
	Health  contracts.CandidateHealth
}

type Orchestrator struct {
	Scorer CandidateScorer
}

func (o Orchestrator) Rank(inputs []CandidateInput) []contracts.RankedCandidate {
	out := make([]contracts.RankedCandidate, 0, len(inputs))
	for _, in := range inputs {
		score := o.Scorer.Score(in.Health)
		reason := "balanced score"
		if in.Health.HandshakeOK && in.Health.LatencyMs < 350 {
			reason = "fast handshake path"
		}
		out = append(out, contracts.RankedCandidate{Profile: in.Profile, Score: score, Reason: reason})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Score > out[j].Score })
	return out
}
