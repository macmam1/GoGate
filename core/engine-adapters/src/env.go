package engineadapters

import "fmt"

func mapEnv(in map[string]string) []string {
	if len(in) == 0 {
		return nil
	}
	out := make([]string, 0, len(in))
	for k, v := range in {
		out = append(out, fmt.Sprintf("%s=%s", k, v))
	}
	return out
}
