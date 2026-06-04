package engineadapters

import (
	"testing"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

func TestBuildSingboxArgs(t *testing.T) {
	req := contracts.SessionStartRequest{
		RuntimeConfigPath: "/tmp/singbox.json",
		AdvancedModeEnable: true,
		Hints: contracts.RouteHints{EdgeIP: "1.2.3.4"},
	}
	args, err := buildSingboxArgs(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(args) < 3 {
		t.Fatalf("unexpected args length: %d", len(args))
	}
}
