package harvester

import (
	"time"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

func RetryPolicyForMode(mode contracts.Mode) RetryPolicy {
	switch mode {
	case contracts.ModeRestrictedNetwork:
		return RetryPolicy{Attempts: 5, BaseDelay: 500 * time.Millisecond, MaxDelay: 6 * time.Second}
	case contracts.ModeBatterySaver:
		return RetryPolicy{Attempts: 2, BaseDelay: 300 * time.Millisecond, MaxDelay: 2 * time.Second}
	default:
		return defaultRetryPolicy()
	}
}
