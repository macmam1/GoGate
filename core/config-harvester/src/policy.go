package harvester

import "strings"

type SourcePolicy struct {
	AllowedPrefixes []string
}

func (p SourcePolicy) IsAllowed(source string) bool {
	if len(p.AllowedPrefixes) == 0 {
		return false
	}
	for _, pref := range p.AllowedPrefixes {
		if strings.HasPrefix(source, pref) {
			return true
		}
	}
	return false
}
