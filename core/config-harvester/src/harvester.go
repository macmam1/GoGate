package harvester

import (
	"fmt"
	"strings"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type RawEntry struct {
	Source string
	Value  string
}

func Normalize(entries []RawEntry) ([]contracts.NormalizedProfile, error) {
	out := make([]contracts.NormalizedProfile, 0, len(entries))
	seen := map[string]struct{}{}

	for idx, e := range entries {
		parts := strings.Split(e.Value, ":")
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid raw entry at index %d", idx)
		}
		id := strings.TrimSpace(parts[0])
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, contracts.NormalizedProfile{
			ID:       id,
			Source:   e.Source,
			Protocol: strings.TrimSpace(parts[1]),
			Endpoint: strings.TrimSpace(parts[2]),
			Port:     443,
		})
	}
	return out, nil
}
