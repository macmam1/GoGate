package engineadapters

import (
	"testing"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

func TestBuildXrayArgs(t *testing.T) {
	req := contracts.SessionStartRequest{
		RuntimeConfigPath: "/tmp/xray.json",
		ConnectTimeoutMs:  5000,
		AdvancedModeEnable: true,
		Hints: contracts.RouteHints{HostSNI: "example.com"},
	}
	args, err := buildXrayArgs(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(args) < 4 {
		t.Fatalf("unexpected args length: %d", len(args))
	}
}
