package harvester

import "github.com/macmam1/GoGate/core/internal/contracts"

func DedupeByID(in []contracts.NormalizedProfile) []contracts.NormalizedProfile {
	out := make([]contracts.NormalizedProfile, 0, len(in))
	seen := map[string]struct{}{}
	for _, p := range in {
		if _, ok := seen[p.ID]; ok {
			continue
		}
		seen[p.ID] = struct{}{}
		out = append(out, p)
	}
	return out
}
