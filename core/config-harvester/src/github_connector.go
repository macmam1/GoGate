package harvester

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GitHubSource struct {
	Owner string
	Repo  string
	Path  string
	Ref   string
	Token string
}

type githubContentResponse struct {
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
}

func FetchGitHubFile(src GitHubSource) ([]RawEntry, error) {
	ref := src.Ref
	if strings.TrimSpace(ref) == "" {
		ref = "main"
	}
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s", src.Owner, src.Repo, src.Path, ref)
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}
	if src.Token != "" {
		req.Header.Set("Authorization", "Bearer "+src.Token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("github source fetch failed: %s: %s", resp.Status, string(body))
	}

	var parsed githubContentResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	if parsed.Encoding != "base64" {
		return nil, fmt.Errorf("unsupported encoding: %s", parsed.Encoding)
	}
	raw, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(parsed.Content, "\n", ""))
	if err != nil {
		return nil, err
	}

	sourceLabel := fmt.Sprintf("github:%s/%s/%s@%s", src.Owner, src.Repo, src.Path, ref)
	lines := strings.Split(string(raw), "\n")
	out := make([]RawEntry, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		out = append(out, RawEntry{Source: sourceLabel, Value: line})
	}
	return out, nil
}
