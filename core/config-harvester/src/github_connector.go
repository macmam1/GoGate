package harvester

import (
    "context"
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
    return FetchGitHubFileWithClient(context.Background(), src, defaultHTTPClient())
}

func FetchGitHubFileWithClient(ctx context.Context, src GitHubSource, client HTTPDoer) ([]RawEntry, error) {
    return FetchGitHubFileWithOptions(ctx, src, client, FetchOptions{})
}

func FetchGitHubFileWithOptions(ctx context.Context, src GitHubSource, client HTTPDoer, opt FetchOptions) ([]RawEntry, error) {
    cfg := normalizeFetchOptions(opt)
    ref := src.Ref
    if strings.TrimSpace(ref) == "" {
        ref = "main"
    }
    apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s", src.Owner, src.Repo, src.Path, ref)

    resp, err := doRequestWithRetry(ctx, client, func() (*http.Request, error) {
        req, err := http.NewRequest(http.MethodGet, apiURL, nil)
        if err != nil {
            return nil, err
        }
        if src.Token != "" {
            req.Header.Set("Authorization", "Bearer "+src.Token)
        }
        req.Header.Set("Accept", "application/vnd.github+json")
        if cfg.UserAgent != "" {
            req.Header.Set("User-Agent", cfg.UserAgent)
        }
        return req, nil
    }, cfg.Retry)
    if err != nil {
        reportFetch(cfg, apiURL, 0, err)
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        body, _ := io.ReadAll(resp.Body)
        err := fmt.Errorf("github source fetch failed: %s: %s", resp.Status, string(body))
        reportFetch(cfg, apiURL, 0, err)
        return nil, err
    }

    var parsed githubContentResponse
    if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
        if cfg.Metrics != nil {
            cfg.Metrics.RecordParseFail()
        }
        reportFetch(cfg, apiURL, 0, err)
        return nil, err
    }

    if parsed.Encoding != "base64" {
        err := fmt.Errorf("unsupported encoding: %s", parsed.Encoding)
        if cfg.Metrics != nil {
            cfg.Metrics.RecordParseFail()
        }
        reportFetch(cfg, apiURL, 0, err)
        return nil, err
    }
    raw, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(parsed.Content, "\n", ""))
    if err != nil {
        if cfg.Metrics != nil {
            cfg.Metrics.RecordParseFail()
        }
        reportFetch(cfg, apiURL, 0, err)
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
    reportFetch(cfg, apiURL, len(out), nil)
    return out, nil
}
