package harvester

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func FetchSubscription(url string) ([]RawEntry, error) {
	resp, err := http.Get(url) // #nosec G107 -- user-provided source validated at policy layer
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("subscription fetch failed: %s", resp.Status)
	}

	entries := make([]RawEntry, 0)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		entries = append(entries, RawEntry{Source: url, Value: line})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}
