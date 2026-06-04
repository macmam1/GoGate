package orchestrator

import (
	"github.com/macmam1/GoGate/core/internal/contracts"
	probe "github.com/macmam1/GoGate/core/probe-engine/src"
	scoring "github.com/macmam1/GoGate/core/scoring-engine/src"
)

func (o Orchestrator) RankFromProbeResults(profiles []contracts.NormalizedProfile, results []probe.Result) []contracts.RankedCandidate {
	if o.Scorer == nil {
		return nil
	}
	byID := map[string]contracts.NormalizedProfile{}
	for _, p := range profiles {
		byID[p.ID] = p
	}
	inputs := make([]CandidateInput, 0, len(results))
	for _, r := range results {
		profile, ok := byID[r.ProfileID]
		if !ok {
			continue
		}
		health := scoring.HealthFromProbeResult(r)
		inputs = append(inputs, CandidateInput{Profile: profile, Health: health})
	}
	return o.Rank(inputs)
}
