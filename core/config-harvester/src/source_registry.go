package harvester

import (
	"bufio"
	"os"
	"strings"
)

type SourceRegistry struct {
	Policy SourcePolicy
}

func LoadSourceRegistry(path string) (SourceRegistry, error) {
	f, err := os.Open(path)
	if err != nil {
		return SourceRegistry{}, err
	}
	defer f.Close()

	prefixes := make([]string, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		prefixes = append(prefixes, line)
	}
	if err := s.Err(); err != nil {
		return SourceRegistry{}, err
	}
	return SourceRegistry{Policy: SourcePolicy{AllowedPrefixes: prefixes}}, nil
}

func (r SourceRegistry) Allowed(source string) bool {
	return r.Policy.IsAllowed(source)
}
