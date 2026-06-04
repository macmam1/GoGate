package harvester

type FetchOptions struct {
	Retry     RetryPolicy
	UserAgent string
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
