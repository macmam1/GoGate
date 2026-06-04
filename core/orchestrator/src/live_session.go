package orchestrator

import (
	"context"
	"errors"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type SessionExecutor interface {
	Attempt(context.Context, contracts.NormalizedProfile) error
}

func (o Orchestrator) ConnectWithFallback(ctx context.Context, ranked []contracts.RankedCandidate, maxDepth int, exec SessionExecutor) (contracts.RankedCandidate, error) {
	if exec == nil {
		return contracts.RankedCandidate{}, errors.New("nil session executor")
	}
	seq := FallbackPlan{Candidates: ranked, MaxDepth: maxDepth}.Sequence()
	if len(seq) == 0 {
		return contracts.RankedCandidate{}, errors.New("no fallback candidates")
	}

	var lastErr error
	for _, c := range seq {
		if err := exec.Attempt(ctx, c.Profile); err == nil {
			return c, nil
		} else {
			lastErr = err
		}
	}
	if lastErr == nil {
		lastErr = errors.New("all attempts failed")
	}
	return contracts.RankedCandidate{}, lastErr
}
