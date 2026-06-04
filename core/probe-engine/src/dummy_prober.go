package probe

import (
	"context"
	"math/rand"
)

type DummyProber struct{}

func (DummyProber) Probe(_ context.Context, req Request) Result {
	lat := 80 + rand.Intn(120)
	return Result{
		ProfileID:      req.Profile.ID,
		LatencyMs:      lat,
		PacketLossPct:  0,
		HandshakeOK:    true,
		StabilityScore: 0.85,
		Err:            nil,
	}
}
