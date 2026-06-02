package engineadapters

import (
	"fmt"
	"time"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type SingboxAdapter struct {
	sup     *Supervisor
	lastErr error
}

func NewSingboxAdapter(binaryPath string, args []string) *SingboxAdapter {
	return &SingboxAdapter{sup: NewSupervisor(ProcessSpec{Name: binaryPath, Args: args})}
}

func (a *SingboxAdapter) Capabilities() AdapterCapabilities {
	return AdapterCapabilities{EngineID: "sing-box", Protocols: []string{"vmess", "vless", "trojan", "hysteria2", "tuic"}, SupportsRouteOverride: true}
}

func (a *SingboxAdapter) Start(req contracts.SessionStartRequest) (SessionHandle, error) {
	if err := a.sup.Start(); err != nil {
		a.lastErr = err
		return SessionHandle{}, err
	}
	return SessionHandle{ID: fmt.Sprintf("singbox-%d", time.Now().UnixNano()), EngineID: "sing-box"}, nil
}

func (a *SingboxAdapter) Health(handle SessionHandle) (HealthSnapshot, error) {
	if !a.sup.Running() {
		return HealthSnapshot{Status: "stopped"}, nil
	}
	return HealthSnapshot{Status: "running", LatencyMs: 0, HandshakeOK: true}, nil
}

func (a *SingboxAdapter) Stop(handle SessionHandle) error {
	err := a.sup.Stop()
	if err != nil {
		a.lastErr = err
	}
	return err
}

func (a *SingboxAdapter) LastError() error { return a.lastErr }
