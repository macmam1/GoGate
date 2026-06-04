package probe

import (
	"context"
	"testing"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

func TestWorkerPoolRun(t *testing.T) {
	pool := NewWorkerPool(DummyProber{}, 2)
	req := []Request{
		{Profile: contracts.NormalizedProfile{ID: "p1"}, Mode: ModeQuick},
		{Profile: contracts.NormalizedProfile{ID: "p2"}, Mode: ModeQuick},
		{Profile: contracts.NormalizedProfile{ID: "p3"}, Mode: ModeDeep},
	}
	res := pool.Run(context.Background(), req)
	if len(res) != 3 {
		t.Fatalf("expected 3 results, got %d", len(res))
	}
}
