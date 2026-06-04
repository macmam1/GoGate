package harvester

import (
	"testing"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

func TestRetryPolicyForMode(t *testing.T) {
	r := RetryPolicyForMode(contracts.ModeRestrictedNetwork)
	if r.Attempts < 4 {
		t.Fatalf("expected higher retry attempts for restricted mode")
	}

	b := RetryPolicyForMode(contracts.ModeBatterySaver)
	if b.Attempts > r.Attempts {
		t.Fatalf("expected battery saver to have fewer retries")
	}
}
