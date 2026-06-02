package engineadapters

import "github.com/macmam1/GoGate/core/internal/contracts"

type AdapterCapabilities struct {
	EngineID              string
	Protocols             []string
	SupportsRouteOverride bool
}

type SessionHandle struct {
	ID       string
	EngineID string
}

type HealthSnapshot struct {
	Status      string
	LatencyMs   int
	HandshakeOK bool
}

type EngineAdapter interface {
	Capabilities() AdapterCapabilities
	Start(req contracts.SessionStartRequest) (SessionHandle, error)
	Health(handle SessionHandle) (HealthSnapshot, error)
	Stop(handle SessionHandle) error
	LastError() error
}
