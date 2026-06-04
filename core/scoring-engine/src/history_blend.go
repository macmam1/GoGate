package scoring

type HistoryProvider interface {
	RecentSuccess(profileID string) (float64, bool)
}

type BlendConfig struct {
	CurrentWeight float64
	HistoryWeight float64
}

func DefaultBlendConfig() BlendConfig {
	return BlendConfig{CurrentWeight: 0.8, HistoryWeight: 0.2}
}

func BlendScore(current float64, profileID string, provider HistoryProvider, cfg BlendConfig) float64 {
	if provider == nil {
		return current
	}
	history, ok := provider.RecentSuccess(profileID)
	if !ok {
		return current
	}
	cw := cfg.CurrentWeight
	hw := cfg.HistoryWeight
	if cw <= 0 && hw <= 0 {
		cw, hw = 0.8, 0.2
	}
	total := cw + hw
	if total <= 0 {
		return current
	}
	return (cw*current + hw*history) / total
}
