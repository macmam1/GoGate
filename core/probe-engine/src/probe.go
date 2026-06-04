package probe

import (
	"context"
	"time"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type Mode string

const (
	ModeQuick Mode = "quick"
	ModeDeep  Mode = "deep"
)

type Request struct {
	Profile contracts.NormalizedProfile
	Mode    Mode
	Timeout time.Duration
}

type Result struct {
	ProfileID      string
	LatencyMs      int
	PacketLossPct  float64
	HandshakeOK    bool
	StabilityScore float64
	Err            error
}

type Prober interface {
	Probe(context.Context, Request) Result
}
