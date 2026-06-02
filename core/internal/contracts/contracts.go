package contracts

type Mode string

const (
	ModeSafeDefault      Mode = "safe-default"
	ModeRestrictedNetwork Mode = "restricted-network"
	ModeBatterySaver     Mode = "battery-saver"
)

type NormalizedProfile struct {
	ID       string
	Source   string
	Protocol string
	Endpoint string
	Port     int
	SNIHint  string
	Tags     []string
}

type CandidateHealth struct {
	LatencyMs      int
	PacketLossPct  float64
	HandshakeOK    bool
	RecentSuccess  float64
	StabilityScore float64
}

type RankedCandidate struct {
	Profile NormalizedProfile
	Score   float64
	Reason  string
}

type RouteHints struct {
	EdgeIP  string
	HostSNI string
}

type SessionStartRequest struct {
	Profile            NormalizedProfile
	Mode               Mode
	AdvancedModeEnable bool
	Hints              RouteHints
}
