package harvester

type FetchObserver interface {
    OnFetchResult(source string, itemCount int, err error)
}

type FetchOptions struct {
    Retry     RetryPolicy
    UserAgent string
    Observer  FetchObserver
    Metrics   *IngestionMetrics
}

func normalizeFetchOptions(in FetchOptions) FetchOptions {
    if in.Retry.Attempts == 0 {
        in.Retry = defaultRetryPolicy()
    }
    if in.UserAgent == "" {
        in.UserAgent = "GoGate-Harvester/0.1"
    }
    return in
}

func reportFetch(opt FetchOptions, source string, itemCount int, err error) {
    if opt.Observer != nil {
        opt.Observer.OnFetchResult(source, itemCount, err)
    }
    if opt.Metrics != nil {
        if err != nil {
            opt.Metrics.RecordFetchFail()
        } else {
            opt.Metrics.RecordFetchSuccess(itemCount)
        }
    }
}
