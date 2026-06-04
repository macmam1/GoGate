package harvester

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type RetryPolicy struct {
	Attempts  int
	BaseDelay time.Duration
	MaxDelay  time.Duration
}

func defaultHTTPClient() *http.Client {
	return &http.Client{Timeout: 20 * time.Second}
}

func defaultRetryPolicy() RetryPolicy {
	return RetryPolicy{Attempts: 3, BaseDelay: 400 * time.Millisecond, MaxDelay: 4 * time.Second}
}

func doRequestWithRetry(ctx context.Context, client HTTPDoer, makeReq func() (*http.Request, error), p RetryPolicy) (*http.Response, error) {
	if p.Attempts <= 0 {
		p.Attempts = 1
	}
	if p.BaseDelay <= 0 {
		p.BaseDelay = 300 * time.Millisecond
	}
	if p.MaxDelay <= 0 {
		p.MaxDelay = 3 * time.Second
	}

	var lastErr error
	for attempt := 1; attempt <= p.Attempts; attempt++ {
		req, err := makeReq()
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)

		resp, err := client.Do(req)
		if !shouldRetry(resp, err) {
			return resp, err
		}

		if err != nil {
			lastErr = err
		} else if resp != nil {
			lastErr = fmt.Errorf("retryable response: %s", resp.Status)
			drainAndClose(resp.Body)
		}

		if attempt == p.Attempts {
			break
		}

		delay := retryDelay(resp, attempt, p)
		if waitErr := waitWithContext(ctx, delay); waitErr != nil {
			return nil, waitErr
		}
	}

	if lastErr == nil {
		lastErr = errors.New("retry exhausted")
	}
	return nil, lastErr
}

func shouldRetry(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}
	if resp == nil {
		return true
	}
	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusRequestTimeout {
		return true
	}
	if resp.StatusCode >= 500 {
		return true
	}
	if resp.StatusCode == http.StatusForbidden {
		if strings.TrimSpace(resp.Header.Get("X-RateLimit-Remaining")) == "0" {
			return true
		}
	}
	return false
}

func retryDelay(resp *http.Response, attempt int, p RetryPolicy) time.Duration {
	if resp != nil {
		if d := parseRetryAfter(resp.Header.Get("Retry-After")); d > 0 {
			return clampDuration(d, p.MaxDelay)
		}
		if d := parseRateLimitReset(resp.Header.Get("X-RateLimit-Reset")); d > 0 {
			return clampDuration(d, p.MaxDelay)
		}
	}
	exp := float64(p.BaseDelay) * math.Pow(2, float64(attempt-1))
	return clampDuration(time.Duration(exp), p.MaxDelay)
}

func parseRetryAfter(v string) time.Duration {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0
	}
	if sec, err := strconv.Atoi(v); err == nil {
		if sec < 0 {
			return 0
		}
		return time.Duration(sec) * time.Second
	}
	if t, err := http.ParseTime(v); err == nil {
		d := time.Until(t)
		if d > 0 {
			return d
		}
	}
	return 0
}

func parseRateLimitReset(v string) time.Duration {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0
	}
	sec, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	resetAt := time.Unix(sec, 0)
	d := time.Until(resetAt)
	if d <= 0 {
		return 0
	}
	return d
}

func clampDuration(v, max time.Duration) time.Duration {
	if v <= 0 {
		return 100 * time.Millisecond
	}
	if v > max {
		return max
	}
	return v
}

func waitWithContext(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

func drainAndClose(rc io.ReadCloser) {
	if rc == nil {
		return
	}
	_, _ = io.Copy(io.Discard, rc)
	_ = rc.Close()
}
