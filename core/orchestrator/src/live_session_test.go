package orchestrator

import (
	"context"
	"errors"
	"testing"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type fakeExec struct {
	failUntil int
	calls     int
}

func (f *fakeExec) Attempt(_ context.Context, _ contracts.NormalizedProfile) error {
	f.calls++
	if f.calls <= f.failUntil {
		return errors.New("failed")
	}
	return nil
}

func TestConnectWithFallback(t *testing.T) {
	o := Orchestrator{}
	ranked := []contracts.RankedCandidate{
		{Profile: contracts.NormalizedProfile{ID: "p1"}},
		{Profile: contracts.NormalizedProfile{ID: "p2"}},
	}
	exec := &fakeExec{failUntil: 1}
	winner, err := o.ConnectWithFallback(context.Background(), ranked, 2, exec)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if winner.Profile.ID != "p2" {
		t.Fatalf("expected fallback candidate p2, got %s", winner.Profile.ID)
	}
}
