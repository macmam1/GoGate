package orchestrator

import "github.com/macmam1/GoGate/core/internal/contracts"

type FallbackPlan struct {
	Candidates []contracts.RankedCandidate
	MaxDepth   int
}

func (p FallbackPlan) Sequence() []contracts.RankedCandidate {
	depth := p.MaxDepth
	if depth <= 0 || depth > len(p.Candidates) {
		depth = len(p.Candidates)
	}
	out := make([]contracts.RankedCandidate, 0, depth)
	for i := 0; i < depth; i++ {
		out = append(out, p.Candidates[i])
	}
	return out
}
