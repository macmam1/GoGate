package harvester

import (
    "bufio"
    "context"
    "fmt"
    "net/http"
    "strings"
)

func FetchSubscription(url string) ([]RawEntry, error) {
    return FetchSubscriptionWithClient(context.Background(), url, defaultHTTPClient())
}

func FetchSubscriptionWithClient(ctx context.Context, url string, client HTTPDoer) ([]RawEntry, error) {
    return FetchSubscriptionWithOptions(ctx, url, client, FetchOptions{})
}

func FetchSubscriptionWithOptions(ctx context.Context, url string, client HTTPDoer, opt FetchOptions) ([]RawEntry, error) {
    cfg := normalizeFetchOptions(opt)
    resp, err := doRequestWithRetry(ctx, client, func() (*http.Request, error) {
        req, err := http.NewRequest(http.MethodGet, url, nil)
        if err != nil {
            return nil, err
        }
        if cfg.UserAgent != "" {
            req.Header.Set("User-Agent", cfg.UserAgent)
        }
        return req, nil
    }, cfg.Retry)
    if err != nil {
        reportFetch(cfg, url, 0, err)
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        err := fmt.Errorf("subscription fetch failed: %s", resp.Status)
        reportFetch(cfg, url, 0, err)
        return nil, err
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
        if cfg.Metrics != nil {
            cfg.Metrics.RecordParseFail()
        }
        reportFetch(cfg, url, 0, err)
        return nil, err
    }
    reportFetch(cfg, url, len(entries), nil)
    return entries, nil
}
