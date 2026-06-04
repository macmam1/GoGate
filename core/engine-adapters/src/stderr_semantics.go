package engineadapters

import "strings"

type semanticPattern struct {
	Kind     ErrorKind
	Patterns []string
}

var stderrSemanticPatterns = []semanticPattern{
	{Kind: ErrConfigInvalid, Patterns: []string{"invalid config", "failed to parse", "json: cannot", "yaml: unmarshal errors", "unknown field"}},
	{Kind: ErrPortInUse, Patterns: []string{"address already in use", "failed to bind", "listen tcp", "bind:"}},
	{Kind: ErrPermission, Patterns: []string{"permission denied", "operation not permitted", "access is denied"}},
	{Kind: ErrNetworkDenied, Patterns: []string{"network unreachable", "no route to host", "connection refused", "forbidden", "tls handshake timeout"}},
}

func semanticKindFromStderr(stderr string) ErrorKind {
	lower := strings.ToLower(stderr)
	if strings.TrimSpace(lower) == "" {
		return ErrUnknown
	}
	for _, entry := range stderrSemanticPatterns {
		for _, p := range entry.Patterns {
			if strings.Contains(lower, p) {
				return entry.Kind
			}
		}
	}
	return ErrUnknown
}
